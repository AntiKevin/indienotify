# IndieNotify

![GitHub repo size](https://img.shields.io/github/repo-size/antikevin/indienotify?style=for-the-badge)
![GitHub language count](https://img.shields.io/github/languages/count/antikevin/indienotify?style=for-the-badge)
![GitHub forks](https://img.shields.io/github/forks/antikevin/indienotify?style=for-the-badge)

> Um servidor de notificações e fila de notificações independente e opensource

### Ajustes e melhorias

O projeto ainda está em desenvolvimento e as próximas atualizações serão voltadas para as seguintes tarefas:

- [x] Endpoint de HealthCheck
- [x] Endpoint Rest de envio de notificaçao (PUB/SUB)
- [x] Modulo de envio de notificação (PUB/SUB)
- [ ] Modulo de envio de notificaçao persistente (REDIS-STREAM)
- [ ] Endpoint para Subscribe em um canal de notificação

## 💻 Pré-requisitos

Antes de começar, verifique se você atendeu aos seguintes requisitos:

- Você instalou a versão mais recente da linguagem `Go`
- Você tem uma máquina `Windows, Linux ou Mac>`.

## ☕ Desenvolvendo para o IndieNotify 

Para usar o IndieNotify, siga estas etapas: </br>

dentro do diretório raiz do projeto
```
go install
```

depois:

```
go run cmd/server/main.go
```

Isso 

## ☕ Usando o IndieNotify

Para usar o IndieNotify, siga estas etapas:


`faça uma requisiçao/conexão a qualquer um dos endpoints uteis para você`

## 📫 Contribuindo para o IndieNotify

Para contribuir com o IndieNotify, siga estas etapas:

1. Bifurque este repositório.
2. Crie um branch: `git checkout -b <nome_branch>`.
3. Faça suas alterações e confirme-as: `git commit -m '<mensagem_commit>'` 
4. utilize o conventional commits e mantenha os commits no padrão do repositório
5. Envie para o branch original: `git push origin <nome_do_projeto> / <local>`
6. Crie a solicitação de pull.

Como alternativa, consulte a documentação do GitHub em [como criar uma solicitação pull](https://help.github.com/en/github/collaborating-with-issues-and-pull-requests/creating-a-pull-request).

## 🤝 Mantenedor(es)

<table>
  <tr>
    <td align="center">
      <a href="#" title="defina o título do link">
        <img src="https://avatars3.githubusercontent.com/u/51024849" width="100px;" alt="Foto do Iuri Silva no GitHub"/><br>
        <sub>
          <b>Kevin Rodrigues</b>
        </sub>
      </a>
    </td>
  </tr>
</table>

<!-- ## 😄 Seja um dos contribuidores

Quer fazer parte desse projeto? Clique [AQUI](CONTRIBUTING.md) e leia como contribuir. -->

## 📝 Licença

Esse projeto está sob licença. Veja o arquivo [LICENÇA](LICENSE.md) para mais detalhes.