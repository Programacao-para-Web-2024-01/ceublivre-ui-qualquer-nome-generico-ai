document.getElementById('criarProdutoForm').addEventListener('submit', async function(e) {
    e.preventDefault();

    const nome = document.getElementById('produtoNome').value;
    const quantidade = parseInt(document.getElementById('produtoQuantidade').value);
    const preco = parseFloat(document.getElementById('produtoPreco').value);

    const produto = { nome: nome, quantidade: quantidade, preco: preco };

    try {
        const response = await fetch('http://localhost:8000/produtos', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify(produto)
        });

        if (!response.ok) {
            const errorText = await response.text();
            throw new Error(`Erro ao adicionar produto: ${errorText}`);
        }

        const produtoAdicionado = await response.json();
        console.log('Produto adicionado:', produtoAdicionado);
        document.getElementById('criarProdutoForm').reset();
        alert('Produto adicionado com sucesso!');
    } catch (error) {
        console.error(error);
        alert(error.message);
    }
});

document.getElementById('criarPedidoForm').addEventListener('submit', async function(e) {
    e.preventDefault();

    const cliente = document.getElementById('cliente').value;
    const produtos = document.getElementById('produtos').value.split(',').map(prod => prod.trim());
    const valorPedido = parseFloat(document.getElementById('valorPedido').value);
    const valorFrete = parseFloat(document.getElementById('valorFrete').value);
    const endereco = document.getElementById('endereco').value;

    const pedido = {
        cliente: cliente,
        produtos: produtos,
        valor_pedido: valorPedido,
        valor_frete: valorFrete,
        endereco: endereco
    };

    try {
        const response = await fetch('http://localhost:8000/pedidos', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify(pedido)
        });

        if (!response.ok) {
            const errorText = await response.text();
            throw new Error(`Erro ao criar pedido: ${errorText}`);
        }

        const pedidoCriado = await response.json();
        console.log('Pedido criado:', pedidoCriado);
        document.getElementById('criarPedidoForm').reset();
        alert('Pedido criado com sucesso!');
    } catch (error) {
        console.error(error);
        alert(error.message);
    }
});

document.getElementById('atualizarStatusForm').addEventListener('submit', async function(e) {
    e.preventDefault();

    const pedidoId = document.getElementById('pedidoIdStatus').value;
    const status = document.getElementById('status').value;

    try {
        const response = await fetch(`http://localhost:8000/pedidos/${pedidoId}/status`, {
            method: 'PUT',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify({ status: status })
        });

        if (!response.ok) {
            const errorText = await response.text();
            throw new Error(`Erro ao atualizar status do pedido: ${errorText}`);
        }

        let pedidoAtualizado;
        try {
            pedidoAtualizado = await response.json();
        } catch (e) {
            pedidoAtualizado = null;
        }

        console.log('Status do pedido atualizado:', pedidoAtualizado);
        document.getElementById('atualizarStatusForm').reset();
        alert('Status do pedido atualizado com sucesso!');
    } catch (error) {
        console.error(error);
        alert(error.message);
    }
});


document.getElementById('atualizarEnderecoForm').addEventListener('submit', async function(e) {
    e.preventDefault();

    const pedidoId = document.getElementById('pedidoIdEndereco').value;
    const novoEndereco = document.getElementById('novoEndereco').value;

    try {
        const response = await fetch(`http://localhost:8000/pedidos/${pedidoId}/endereco`, {
            method: 'PUT',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify({ endereco: novoEndereco })
        });

        if (!response.ok) {
            const errorText = await response.text();
            throw new Error(`Erro ao atualizar endereço do pedido: ${errorText}`);
        }

        const pedidoAtualizado = await response.json();
        console.log('Endereço do pedido atualizado:', pedidoAtualizado);
        document.getElementById('atualizarEnderecoForm').reset();
        alert('Endereço do pedido atualizado com sucesso!');
    } catch (error) {
        console.error(error);
        alert(error.message);
    }
});

document.getElementById('rastrearPedidoForm').addEventListener('submit', async function(e) {
    e.preventDefault();

    const pedidoId = document.getElementById('pedidoIdRastrear').value;

    try {
        const response = await fetch(`http://localhost:8000/pedidos/${pedidoId}`);
        if (!response.ok) {
            const errorText = await response.text();
            throw new Error(`Erro ao rastrear pedido: ${errorText}`);
        }

        const pedido = await response.json();
        const detalhesDiv = document.getElementById('pedidoDetalhes');
        detalhesDiv.innerHTML = '';

    
        const pedidoSimp = {
            endereco: pedido.endereco,
            status: pedido.status
        };

        detalhesDiv.innerHTML = `
            <h3>Detalhes do Pedido</h3>
            <p>Endereço: ${pedidoSimp.endereco}</p>
            <p>Status: ${pedidoSimp.status}</p>
        `;
    } catch (error) {
        console.error(error);
        alert(error.message);
    }
});


document.getElementById('listarProdutosBtn').addEventListener('click', async function() {
    try {
        const response = await fetch('http://localhost:8000/produtos');
        if (!response.ok) {
            const errorText = await response.text();
            throw new Error(`Erro ao listar produtos: ${errorText}`);
        }

        const produtos = await response.json();
        const produtosList = document.getElementById('produtosList');
        produtosList.innerHTML = produtos.map(produto => `
            <li>
                <strong>ID:</strong> ${produto.id} <br>
                <strong>Nome:</strong> ${produto.nome} <br>
                <strong>Quantidade:</strong> ${produto.quantidade} <br>
                <strong>Preço:</strong> ${produto.preco}
            </li>
        `).join('');
    } catch (error) {
        console.error(error);
        alert(error.message);
    }
});

document.getElementById('listarPedidosBtn').addEventListener('click', async function() {
    try {
        const response = await fetch('http://localhost:8000/pedidos');
        if (!response.ok) {
            const errorText = await response.text();
            throw new Error(`Erro ao listar pedidos: ${errorText}`);
        }

        const pedidos = await response.json();
        const pedidosList = document.getElementById('pedidosList');
        pedidosList.innerHTML = pedidos.map(pedido => `
            <li>
                <strong>ID:</strong> ${pedido.id} <br>
                <strong>Cliente:</strong> ${pedido.cliente} <br>
                <strong>Produtos:</strong> ${pedido.produtos.join(', ')} <br>
                <strong>Valor do Pedido:</strong> ${pedido.valor_pedido} <br>
                <strong>Valor do Frete:</strong> ${pedido.valor_frete} <br>
                <strong>Endereço:</strong> ${pedido.endereco} <br>
                <strong>Status:</strong> ${pedido.status}
            </li>
        `).join('');
    } catch (error) {
        console.error(error);
        alert(error.message);
    }
});
