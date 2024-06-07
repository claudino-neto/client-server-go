**client.go**

-criei uma lista de string reply(n sabia se aresposta sdo http get é uma lista ou n, mas pelo codigo do get é so uma string)
-criei também uma variavel request pra cliente poder digitar qual requisição que ele quer

-pergunto qual requisição e recebo do cliente

-o dial conecta com aquele servido criado em server, se der erro entra no if err

-defer func ainda n entendi mas acho q é para fechar a conexão do cliente

-args acabou que n usei ent pode apagar

-cleint call é para chamar o método do web, no caso a nóssa lógica em relação ao HTTP e ele recebe a resposta(reply)

**server.go**

-cria um server implemetando a lógica do arquivo de http e registra o nome("HTTP")
-o listen identifica a rede local

-o defer tem q revisar mas é para fechar a conexão

-e o server acceita a rede local para entrar(possibilitando o cliente

**web.go**
-só fiz o http get, perguntei ao gpt como fazi e pra mim fez sentido, a testar
