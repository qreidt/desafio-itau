# Desafio Itaú
Desafio Itaú em várias linguagens com o objetivo de se divertir e aprender

## Requisitos
1. A base de dados local deve conter usuários emissores e receptores de uma transferência, seguindo o modelo.
2. A senha de autenticação deve conter 8 dígitos, caracteres especiais, letras maiúsculas e minúsculas
3. O usuário precisará estar autenticado para realizar a transferência
4. Toda transferência deverá ser feita entre um emissor e um receptor
5. Operador e receptor iniciam a operação com um saldo de $ 0,00
6. As transferências deverão ser executadas de acordo com o seu tipo, sendo 3 os tipos: PIX, TED e DOC
7. O limite de valor máximo permitido para uma transferência via PIX é de $ 5 mil
8. Transferências via TED só são permitidas para valores acima de $ 5 mil e até $ 10 mil
9. Transferências via DOC só são permitidas para valores acima de $ 10 mil
10. Não serão permitidas transferências para a mesma conta
11. Um emissor pode transferir para ele mesmo se for para uma conta diferente
12. As entradas deverão sempre estar com todos os dados preenchidos
13. Se a transferência for bem sucedida, exibir a mensagem de sucesso com o saldo do emissor e do receptor após a transferência, de acordo com o modelo
  ```
    Modelo de mensagem de sucesso:
    Sua transferência foi realizada com sucesso!
      Saldo do emissor: $ X,XX
      Saldo do receptor: $ X,XX
  ```
14. Se a transferência não foi completa, exibir a mensagem de erro explicando o motivo. Ex: "Sua transferência não foi completada pois <motivo>"

## Melhorias propostas além que já foi definido
1. Senha especial para realizar transações
2. Criptografar senhas armazenadas
3. Autenticação com multiplos dígitos de senha
