<div align="center">
<img src="./docs/static/statusphere-white.png" width="300" height="300" alt="Statusphere logo">
</div>

<div align="center">

![GitHub stars](https://img.shields.io/github/stars/metoro-io/statusphere?style=social)
![GitHub forks](https://img.shields.io/github/forks/metoro-io/statusphere?style=social)
![GitHub issues](https://img.shields.io/github/issues/metoro-io/statusphere)
![GitHub pull requests](https://img.shields.io/github/issues-pr/metoro-io/statusphere)
![GitHub license](https://img.shields.io/github/license/metoro-io/statusphere)
![GitHub contributors](https://img.shields.io/github/contributors/metoro-io/statusphere)
![GitHub last commit](https://img.shields.io/github/last-commit/metoro-io/statusphere)

</div>

An open-source api-first status page aggregator.

If your looking for the official hosted version, head [here](https://metoro.io/statusphere).

## Architecture

Statusphere is made up of 3 main components:

1. The scrapers
2. The database
3. The api servers

They're orchestrated in the following way:

<div align="center">
<img src="./docs/static/statusphere-architecture-white.png" height="300" alt="Statusphere logo">
</div>

## Usage

Warning: This will spin up a local instance of the statusphere stack which will automatically scrape the status pages of
the services listed in the `status_pages.go` file.

```bash
docker-compose up

# The api server will be available at http://localhost:8080
curl http://localhost:8080/api/v1/statusPages/count
```
