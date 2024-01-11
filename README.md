<!-- Improved compatibility of back to top link: See: https://github.com/othneildrew/Best-README-Template/pull/73 -->
<a name="readme-top"></a>

<!-- PROJECT SHIELDS -->
[![Contributors][contributors-shield]][contributors-url]
[![Forks][forks-shield]][forks-url]
[![Stargazers][stars-shield]][stars-url]
[![Issues][issues-shield]][issues-url]
[![MIT License][license-shield]][license-url]
[![LinkedIn][linkedin-shield]][linkedin-url]



<!-- PROJECT LOGO -->
<br />
<div align="center">
  <a href="https://github.com/github_username/repo_name">
    <img src="images/logo.png" alt="Logo" width="80" height="80">
  </a>

<h3 align="center">World Holidays</h3>

  <p align="center">
    Retrieve the holidays from a specific country.
    <a href="https://github.com/github_username/repo_name"><strong>Available countries »</strong></a>
    <br />
    <a href="https://github.com/github_username/repo_name"><strong>Explore the docs »</strong></a>
    <br />
    <br />
    <a href="https://github.com/github_username/repo_name/issues">Report Bug</a>
    <a href="https://github.com/github_username/repo_name/issues">Request Feature</a>
  </p>
</div>



<!-- TABLE OF CONTENTS -->
<details>
  <summary>Table of Contents</summary>
  <ol>
    <li>
      <a href="#about-the-project">About The Project</a>
      <ul>
        <li><a href="#built-with">Built With</a></li>
      </ul>
    </li>
    <li>
      <a href="#getting-started">Getting Started</a>
      <ul>
        <li><a href="#prerequisites">Prerequisites</a></li>
        <li><a href="#installation">Installation</a></li>
      </ul>
    </li>
    <li><a href="#usage">Usage</a></li>
    <li><a href="#roadmap">Roadmap</a></li>
    <li><a href="#contributing">Contributing</a></li>
    <li><a href="#license">License</a></li>
    <li><a href="#contact">Contact</a></li>
    <li><a href="#acknowledgments">Acknowledgments</a></li>
  </ol>
</details>



<!-- ABOUT THE PROJECT -->
## About The Project

[![Product Name Screen Shot][product-screenshot]](https://example.com)

Here's a blank template to get started: To avoid retyping too much info. Do a search and replace with your text editor for the following: `github_username`, `repo_name`, `twitter_handle`, `linkedin_username`, `email_client`, `email`, `project_title`, `project_description`

<p align="right">(<a href="#readme-top">back to top</a>)</p>



### Built With

[![Go][Go-url]][Go-url]

<p align="right">(<a href="#readme-top">back to top</a>)</p>



<!-- GETTING STARTED -->
## Getting Started

This is an example of how you may give instructions on setting up your project locally.
To get a local copy up and running follow these simple example steps.

### Prerequisites

This is an example of how to list things you need to use the software and how to install them.
* npm
  ```sh
  npm install npm@latest -g
  ```

### Installation

1. Get a free API Key at [https://example.com](https://example.com)
2. Clone the repo
   ```sh
   git clone https://github.com/github_username/repo_name.git
   ```
3. Install NPM packages
   ```sh
   npm install
   ```
4. Enter your API in `config.js`
   ```js
   const API_KEY = 'ENTER YOUR API';
   ```

<p align="right">(<a href="#readme-top">back to top</a>)</p>



<!-- MARKDOWN LINKS & IMAGES -->
<!-- https://www.markdownguide.org/basic-syntax/#reference-style-links -->
[contributors-shield]: https://img.shields.io/github/contributors/github_username/repo_name.svg?style=for-the-badge
[contributors-url]: https://github.com/github_username/repo_name/graphs/contributors
[forks-shield]: https://img.shields.io/github/forks/github_username/repo_name.svg?style=for-the-badge
[forks-url]: https://github.com/github_username/repo_name/network/members
[stars-shield]: https://img.shields.io/github/stars/github_username/repo_name.svg?style=for-the-badge
[stars-url]: https://github.com/github_username/repo_name/stargazers
[issues-shield]: https://img.shields.io/github/issues/github_username/repo_name.svg?style=for-the-badge
[issues-url]: https://github.com/github_username/repo_name/issues
[license-shield]: https://img.shields.io/github/license/github_username/repo_name.svg?style=for-the-badge
[license-url]: https://github.com/github_username/repo_name/blob/master/LICENSE.txt
[linkedin-shield]: https://img.shields.io/badge/-LinkedIn-black.svg?style=for-the-badge&logo=linkedin&colorB=555
[linkedin-url]: https://linkedin.com/in/linkedin_username
[product-screenshot]: images/screenshot.png
[Go-url]: https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white





# Idea

Crear una API REST con GO que devuelva los feriados para un pais especifico.

Como?

1) Crear una base de datos con un modelo del tipo descrito debajo. Agregar una API, usar Gingonic para crear un endpoint. Que, cuando alguien consulte a la API, busque holidays
para el pais recibido en la URL dentro de la base, filtrando el año solicitado si ningun feriado existe, disparar un Script que complete los datos y retornar la respuesta.
El script debería pegarle a https://date.nager.at/api/v3/publicholidays/2024/AR, responder al usuario y correr un metodo que almacene los feriados nuevos.

La arquitectura: Lambdas? EC2 t2 micro que autoescalen horizontalmente, que esten dentro de una red privada conectadas a una tabla no relacional. Load balancer derivando las requests. El Load balancer tiene acceso a la red privada de EC2s y tiene un endpoint y puerto publico. Route 53 con un alias al load balancer.

+ Performance
+ Codigo

2) No crear una base. Actuar como interfaz entre la API date.nager. Crear una API que tenga un
metodo que le pegue a la API y retorne la respuesta.

La arquitectura, la misma sin la Base de datos.

* Country: codigo(PK), nombre, estado(activo o no).
* Holiday: fecha, pais(codigo PK), tipo de feriado


Endpoint URL: api/v1/publicholidays/2024/AR
Endpoint Response:

    ```
    [
    {
        "date": "2024-01-01",
        "localName": "Año Nuevo",
        "name": "New Year's Day",
        "countryCode": "AR",
        "fixed": true,
        "global": true,
        "counties": null,
        "launchYear": null,
        "types": [
            "Public"
        ]
    },
    {
        "date": "2024-02-12",
        "localName": "Carnaval",
        "name": "Carnival",
        "countryCode": "AR",
        "fixed": false,
        "global": true,
        "counties": null,
        "launchYear": null,
        "types": [
            "Public"
        ]
    }
    ]
    ```