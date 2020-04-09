# Guia de contribuição do projeto

## Tabela de conteúdo

- [Primeiros Passos](#primeiros-passos)
  - [Código de conduta](#código-de-conduta)
- [Como posso ajudar?](#como-posso-ajudar)
  - [Documentação](#documentação)
  - [Problemas](#problemas)
    - [Enviando um problema](#enviando-um-problema)
  - [Feedback](#feedback)
  - [Código](#código)
    - [Ambiente de desenvolvimento](#ambiente-de-desenvolvimento)
- [Convenção de Commits](#convenção-de-commits)
  - [Por que todas essas regras?](#porque-todas-essas-regras)
- [Enviando um pull request](#enviando-um-pull-request)

## Primeiros Passos

Primeiramente, gostaríamos de agradecer por você dedicar um tempo para contribuir e fazer deste um projeto melhor.

Aqui temos um conjunto de instruções e diretrizes para reduzir mal-entendidos e tornar o processo de contribuição para o `codeedu-plataforma-desafios` o mais suave possível.

Esperamos que este guia ajude a torne claro o processo de contribuição e responda a quaisquer perguntas que você possa ter.

### Código de conduta

Esperamos que os participantes do projeto cumpram nosso Código de Conduta. Você pode verificar o [texto completo](https://opensource.guide/code-of-conduct/) para entender o tipo de conduta que esperamos e quais ações serão e não serão toleradas.

Ao participar deste projeto, você concorda em cumprir seus termos.

## Como posso ajudar?

Aqui estão algumas maneiras pelas quais você pode ajudar, juntamente com algumas diretrizes.

### Documentação

Como usuário de `codeedu-plataforma-desafios`, você é o candidato perfeito para nos ajudar a melhorar nossa documentação.

Erros de digitação, erros, falta de exemplos e/ou explicações e assim por diante, são apenas alguns exemplos de coisas que podem ser corrigidas e/ou aprimoradas.

Você pode até fazer melhorias neste guia. :)

Ao documentar, tente manter as coisas simples e claras.

### Problemas

Alguns problemas são criados com informações ausentes, sem um modelo, não reproduzíveis ou simples
inválido.

Você pode torná-los mais fáceis de entender e resolver.

#### Enviando um problema

- Procure problemas semelhantes antes de abrir um novo;
- Use um dos modelos de problemas correspondentes;
- Use um título claro e descritivo;
- Inclua o máximo de informações possível, preenchendo o modelo de problema fornecido;
- Na maioria das vezes, a melhor maneira de relatar um problema é a falha no teste.

### Feedback

Quanto mais feedback, melhor! Estamos sempre procurando mais sugestões e opiniões sobre discussões. Essa é uma boa oportunidade para influenciar a direção futura desta ferramenta.

Isso inclui o envio de uma sugestão de aprimoramento, incluindo recursos completamente novos e pequenas melhorias na funcionalidade existente.

### Código

Você pode usar `issue labels` para descobrir problemas com os quais você poderia ajudar:

- [`bug` issues](https://github.com/codeedu/codeedu-plataforma-desafios/labels/bug) são erros conhecidos que gostaríamos de corrigir;
- [`enhancement` issues](https://github.com/codeedu/codeedu-plataforma-desafios/labels/enhancement) são recursos que estamos abertos a incluir.

Quando você vir um problema(`issue`) que já está atribuído, verifique se já não há alguém trabalhando nele (talvez tente perguntar no problema(`issue`)). Isso é para evitar trabalho desnecessário para todos os envolvidos.

#### Ambiente de desenvolvimento

Ao desenvolver, prefira usar **Go** ≥ 1,14. Escrever código com as últimas versões estáveis ​​do Go nos permite usar novas ferramentas de desenvolvedor.

Depois de [clonar o repositório](https://help.github.com/articles/cloning-a-repository/), para subir as aplicações `client` e `server` você precisa navegar até suas respectivas pastas.

**client** (/codeedu-plataforma-desafios/framework/cmd/client)

    $ go run main.go

**server** (/codeedu-plataforma-desafios/framework/cmd/server)

    $ go run main.go

**Despendências:**
Você precisa subir a base de dados usando o `docker` para isso, basta execultar o comando abaixo.

    $ docker-compose up

#### Convenção de Commits

Essa convenção usa o [conventionalcommits.org](https://www.conventionalcommits.org/en/v1.0.0/)

#### Resumo

A especificação do Conventional Commits é uma convenção simples para utilizar nas mensagens de commit. Ela define um conjunto de regras para criar um histórico de commit explícito, o que facilita a criação de ferramentas automatizadas. Esta convenção segue o [SemVer](https://semver.org/), descrevendo os recursos, correções e modificações que quebram a compatibilidade nas mensagens de commit.

A mensagem do commit deve ser estruturada da seguinte forma:

```
<type>[optional scope]: <description>

[optional body]

[optional footer(s)]
```

O commit contém os seguintes elementos estruturais, para comunicar a intenção ao utilizador da sua biblioteca:

**fix:** um commit do tipo fix soluciona um problema na sua base de código (isso se correlaciona com `PATCH` do versionamento semântico).
**feat:** um commit do tipo feat inclui um novo recurso na sua base de código (isso se correlaciona com `MINOR` do versionamento semântico).
**BREAKING CHANGE:** um commit que contém o texto BREAKING CHANGE:, no começo do texto do corpo opcional ou do rodapé opcional, inclui uma modificação que quebra a compatibilidade da API (isso se correlaciona com `MAJOR` do versionamento semântico). Uma BREAKING CHANGE pode fazer parte de commits de qualquer tipo.
Outros: tipos adicionais são permitidos além de fix: e feat:, por exemplo [@commitlint/config-conventional](https://github.com/conventional-changelog/commitlint/tree/master/%40commitlint/config-conventional) (baseado na Convenção do Angular) recomenda-se `chore:`, `docs:`, `style:`, `refactor:`, `perf:`, `test:`, entre outros.

##### [Exemplos de commits](https://www.conventionalcommits.org/pt-br/v1.0.0-beta.4/#exemplos)

##### [mais informações](https://www.conventionalcommits.org/en/v1.0.0/)

## Enviando um pull request

Antes de enviar um `pull request`, verifique se o seguinte foi feito:

- Está segguindo o [GitHub Flow](https://guides.github.com/introduction/flow/)
- Está segguindo o [5 Useful Tips For A Better Commit Message](https://robots.thoughtbot.com/5-useful-tips-for-a-better-commit-message) artigo e o post [How to Write a Git Commit Message](http://chris.beams.io/posts/git-commit/) post.
- Está segguindo o [Linking a pull request to an issue](https://help.github.com/en/github/managing-your-work-on-github/linking-a-pull-request-to-an-issue)

- [Fork](https://help.github.com/en/articles/fork-a-repo) o repositório e crie seu branch a partir do `master`.
  - Exemplo: `feature/my-awesome-feature` ou `fix/annoying-bug`;
- Execute `go run main.go` no `client` e no `server`;
- Se você corrigiu um bug ou adicionou um código que deve ser testado, **adicione testes** por favor;
- Garantir que o conjunto de testes seja aprovado;
- Garanta que seu commit seja validado;
