<h1 align="center"><img src="https://raw.githubusercontent.com/ory/meta/master/static/banners/keto.svg" alt="Ory Keto - Open Source & Cloud Native Access Control Server"></h1>

<h4 align="center">    
    <a href="https://www.ory.sh/chat">Chat</a> |
    <a href="https://github.com/ory/keto/discussions">Discusssions</a> |
    <a href="https://www.ory.sh/l/sign-up-newsletter">Newsletter</a><br/><br/>
    <a href="https://www.ory.sh/docs/keto/">Guide</a> |
    <a href="https://www.ory.sh/docs/keto/sdk/api">API Docs</a> |
    <a href="https://godoc.org/github.com/ory/keto">Code Docs</a><br/><br/>
    <a href="https://console.ory.sh/">Support this project!</a><br/><br/>
    <a href="https://www.ory.sh/jobs/">Work in Open Source, Ory is hiring!</a>
</h4>

---

<p align="left">
    <a href="https://github.com/ory/keto/actions/workflows/ci.yaml"><img src="https://github.com/ory/keto/actions/workflows/ci.yaml/badge.svg?branch=master&event=push" alt="CI Tasks for Ory keto"></a>
    <a href="https://coveralls.io/github/ory/keto?branch=master"> <img src="https://coveralls.io/repos/ory/keto/badge.svg?branch=master&service=github" alt="Coverage Status"></a>
    <a href="https://goreportcard.com/report/github.com/ory/keto"><img src="https://goreportcard.com/badge/github.com/ory/keto" alt="Go Report Card"></a>
    <a href="https://pkg.go.dev/github.com/ory/keto"><img src="https://pkg.go.dev/badge/www.github.com/ory/keto" alt="PkgGoDev"></a>
    <a href="#backers" alt="sponsors on Open Collective"><img src="https://opencollective.com/ory/backers/badge.svg" /></a> <a href="#sponsors" alt="Sponsors on Open Collective"><img src="https://opencollective.com/ory/sponsors/badge.svg" /></a>
    <a href="https://github.com/ory/keto/blob/master/CODE_OF_CONDUCT.md" alt="Ory Code of Conduct"><img src="https://img.shields.io/badge/ory-code%20of%20conduct-green" /></a>
</p>

Ory Keto is the first and most popular open source implementation of "Zanzibar:
Google's Consistent, Global Authorization System"!

## Get Started

You can use
[Docker to run Ory Keto locally](https://www.ory.sh/docs/keto/install) or use
the Ory CLI to try out Ory Keto:

```sh
# This example works best in Bash
bash <(curl https://raw.githubusercontent.com/ory/meta/master/install.sh) -b . ory
sudo mv ./ory /usr/local/bin/

# Or with Homebrew installed
brew install ory/tap/cli
```

create a new project (you may also use
[Docker](https://www.ory.sh/docs/keto/install))

```sh
ory create project --name "Ory Keto Example"
export project_id="{set to the id from output}"
```

and follow the quick & easy steps below.

## Create a namespace with the Ory Permission Language

```sh
# Write a simple configuration with one namespace
echo "class Document implements Namespace {}" > config.ts

# Apply that configuration
ory patch opl --project $project_id -f file://./config.ts

# Create a relationship that grants tom access to a document
echo "Document:secret#read@tom" \
  | ory parse relation-tuples --project=$project_id --format=json - \
  | ory create relation-tuples --project=$project_id -

# List all relationships
ory list relation-tuples --project=$project_id
# Output:
#   NAMESPACE	OBJECT	RELATION NAME	SUBJECT
#   Document	secret	read		tom
#
#   NEXT PAGE TOKEN
#   IS LAST PAGE	true
```

Now, check out your project on the [Ory Network](https://console.ory.sh/) or
continue with a [more in-depth guide](https://www.ory.sh/docs/keto/quickstart).

### Ory Keto on the Ory Network

The [Ory Network](https://www.ory.sh/cloud) is the fastest, most secure and
worry-free way to use Ory's Services. **Ory Permissions** is powered by the Ory
Keto open source permission server, and it's fully API-compatible.

The Ory Network provides the infrastructure for modern end-to-end security:

- Identity & credential management scaling to billions of users and devices
- Registration, Login and Account management flows for passkey, biometric,
  social, SSO and multi-factor authentication
- Pre-built login, registration and account management pages and components
- OAuth2 and OpenID provider for single sign on, API access and
  machine-to-machine authorization
- **Low-latency permission checks based on Google's Zanzibar model and with
  built-in support for the Ory Permission Language**

It's fully managed, highly available, developer & compliance-friendly!

- GDPR-friendly secure storage with data locality
- Cloud-native APIs, compatible with Ory's Open Source servers
- Comprehensive admin tools with the web-based Ory Console and the Ory Command
  Line Interface (CLI)
- Extensive documentation, straightforward examples and easy-to-follow guides
- Fair, usage-based [pricing](https://www.ory.sh/pricing)

Sign up for a
[**free developer account**](https://console.ory.sh/registration?utm_source=github&utm_medium=banner&utm_campaign=keto-readme)
today!

## Ory Keto On-premise support

Are you running Ory Keto in a mission-critical, commercial environment? The Ory
Enterprise License (OEL) provides enhanced features, security, and expert
support directly from the Ory core maintainers.

Organizations that require advanced features, enhanced security, and
enterprise-grade support for Ory's identity and access management solutions
benefit from the Ory Enterprise License (OEL) as a self-hosted, premium offering
including:

- Additional features not available in the open-source version.
- Regular releases that address CVEs and security vulnerabilities, with strict
  SLAs for patching based on severity.
- Support for advanced scaling and multi-tenancy features.
- Premium support options, including SLAs, direct engineer access, and concierge
  onboarding.
- Access to private Docker registry for a faster, more reliable access to vetted
  enterprise builds.

A valid Ory Enterprise License and access to the Ory Enterprise Docker Registry
are required to use these features. OEL is designed for mission-critical,
production, and global applications where organizations need maximum control and
flexibility over their identity infrastructure. Ory's offering is the only
official program for qualified support from the maintainers. For more
information book a meeting with the Ory team to
**[discuss your needs](https://www.ory.sh/contact/)**!

## Ory Permissions, Keto and the Google's Zanzibar model

> Determining whether online users are authorized to access digital objects is
> central to preserving privacy. This paper presents the design, implementation,
> and deployment of Zanzibar, a global system for storing and evaluating access
> control lists. Zanzibar provides a uniform data model and configuration
> language for expressing a wide range of access control policies from hundreds
> of client services at Google, including Calendar, Cloud, Drive, Maps, Photos,
> and YouTube. Its authorization decisions respect causal ordering of user
> actions and thus provide external consistency amid changes to access control
> lists and object contents. Zanzibar scales to trillions of access control
> lists and millions of authorization requests per second to support services
> used by billions of people. It has maintained 95th-percentile latency of less
> than 10 milliseconds and availability of greater than 99.999% over 3 years of
> production use.
>
> [Source](https://research.google/pubs/pub48190/)

If you need to know if a user (or robot, car, service) is allowed to do
something - Ory Permissions and Ory Keto are the right fit for you.

Currently, Ory Permissions [on the Ory Network] and the open-source Ory Keto
permission server implement the API contracts for managing and checking
relations ("permissions") with HTTP and gRPC APIs, as well as global rules
defined through the Ory Permission Language ("userset rewrites"). Future
versions will include features such as Zookies, reverse permission lookups, and
more.

---

<!-- START doctoc generated TOC please keep comment here to allow auto update -->
<!-- DON'T EDIT THIS SECTION, INSTEAD RE-RUN doctoc TO UPDATE -->

- [Ory Permissions, Keto and the Google's Zanzibar model](#ory-permissions-keto-and-the-googles-zanzibar-model)
- [Who's Using It?](#whos-using-it)
  - [Installation](#installation)
- [Ecosystem](#ecosystem)
  - [Ory Kratos: Identity and User Infrastructure and Management](#ory-kratos-identity-and-user-infrastructure-and-management)
  - [Ory Hydra: OAuth2 & OpenID Connect Server](#ory-hydra-oauth2--openid-connect-server)
  - [Ory Oathkeeper: Identity & Access Proxy](#ory-oathkeeper-identity--access-proxy)
  - [Ory Keto: Access Control Policies as a Server](#ory-keto-access-control-policies-as-a-server)
- [Security](#security)
  - [Disclosing Vulnerabilities](#disclosing-vulnerabilities)
- [Telemetry](#telemetry)
  - [Guide](#guide)
  - [HTTP API Documentation](#http-api-documentation)
  - [Upgrading and Changelog](#upgrading-and-changelog)
  - [Command Line Documentation](#command-line-documentation)
  - [Develop](#develop)
    - [Dependencies](#dependencies)
    - [Install From Source](#install-from-source)
    - [Formatting Code](#formatting-code)
    - [Running Tests](#running-tests)
      - [Short Tests](#short-tests)
      - [Regular Tests](#regular-tests)
      - [End-to-End Tests](#end-to-end-tests)
    - [Build Docker](#build-docker)

<!-- END doctoc generated TOC please keep comment here to allow auto update -->

## Who's Using It?

<!--BEGIN ADOPTERS-->

The Ory community stands on the shoulders of individuals, companies, and
maintainers. The Ory team thanks everyone involved - from submitting bug reports
and feature requests, to contributing patches and documentation. The Ory
community counts more than 50.000 members and is growing. The Ory stack protects
7.000.000.000+ API requests every day across thousands of companies. None of
this would have been possible without each and everyone of you!

The following list represents companies that have accompanied us along the way
and that have made outstanding contributions to our ecosystem. _If you think
that your company deserves a spot here, reach out to
<a href="mailto:office@ory.sh">office@ory.sh</a> now_!

<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Logo</th>
            <th>Website</th>
            <th>Case Study</th>
        </tr>
    </thead>
    <tbody>
        <tr>
            <td>OpenAI</td>
            <td align="center">
                <picture>
                    <source media="(prefers-color-scheme: dark)" srcset="https://raw.githubusercontent.com/ory/meta/master/static/adopters/light/openai.svg" />
                    <img height="32px" src="https://raw.githubusercontent.com/ory/meta/master/static/adopters/dark/openai.svg" alt="OpenAI">
                </picture>
            </td>
            <td><a href="https://openai.com/">openai.com</a></td>
            <td><a href="https://www.ory.sh/case-studies/openai">OpenAI Case Study</a></td>
        </tr>
        <tr>
            <td>Fandom</td>
            <td align="center">
                <picture>
                    <source media="(prefers-color-scheme: dark)" srcset="https://raw.githubusercontent.com/ory/meta/master/static/adopters/light/fandom.svg" />
                    <img height="32px" src="https://raw.githubusercontent.com/ory/meta/master/static/adopters/dark/fandom.svg" alt="Fandom">
                </picture>
            </td>
            <td><a href="https://www.fandom.com/">fandom.com</a></td>
            <td><a href="https://www.ory.sh/case-studies/fandom">Fandom Case Study</a></td>
        </tr>
        <tr>
            <td>Lumin</td>
            <td align="center">
                <picture>
                    <source media="(prefers-color-scheme: dark)" srcset="https://raw.githubusercontent.com/ory/meta/master/static/adopters/light/lumin.svg" />
                    <img height="32px" src="https://raw.githubusercontent.com/ory/meta/master/static/adopters/dark/lumin.svg" alt="Lumin">
                </picture>
            </td>
            <td><a href="https://www.luminpdf.com/">luminpdf.com</a></td>
            <td><a href="https://www.ory.sh/case-studies/lumin">Lumin Case Study</a></td>
        </tr>
        <tr>
            <td>Sencrop</td>
            <td align="center">
                <picture>
                    <source media="(prefers-color-scheme: dark)" srcset="https://raw.githubusercontent.com/ory/meta/master/static/adopters/light/sencrop.svg" />
                    <img height="32px" src="https://raw.githubusercontent.com/ory/meta/master/static/adopters/dark/sencrop.svg" alt="Sencrop">
                </picture>
            </td>
            <td><a href="https://sencrop.com/">sencrop.com</a></td>
            <td><a href="https://www.ory.sh/case-studies/sencrop">Sencrop Case Study</a></td>
        </tr>
        <tr>
            <td>OSINT Industries</td>
            <td align="center">
                <picture>
                    <source media="(prefers-color-scheme: dark)" srcset="https://raw.githubusercontent.com/ory/meta/master/static/adopters/light/osint.svg" />
                    <img height="32px" src="https://raw.githubusercontent.com/ory/meta/master/static/adopters/dark/osint.svg" alt="OSINT Industries">
                </picture>
            </td>
            <td><a href="https://www.osint.industries/">osint.industries</a></td>
            <td><a href="https://www.ory.sh/case-studies/osint">OSINT Industries Case Study</a></td>
        </tr>
        <tr>
            <td>HGV</td>
            <td align="center">
                <picture>
                    <source media="(prefers-color-scheme: dark)" srcset="https://raw.githubusercontent.com/ory/meta/master/static/adopters/light/hgv.svg" />
                    <img height="32px" src="https://raw.githubusercontent.com/ory/meta/master/static/adopters/dark/hgv.svg" alt="HGV">
                </picture>
            </td>
            <td><a href="https://www.hgv.it/">hgv.it</a></td>
            <td><a href="https://www.ory.sh/case-studies/hgv">HGV Case Study</a></td>
        </tr>
        <tr>
            <td>Maxroll</td>
            <td align="center">
                <picture>
                    <source media="(prefers-color-scheme: dark)" srcset="https://raw.githubusercontent.com/ory/meta/master/static/adopters/light/maxroll.svg" />
                    <img height="32px" src="https://raw.githubusercontent.com/ory/meta/master/static/adopters/dark/maxroll.svg" alt="Maxroll">
                </picture>
            </td>
            <td><a href="https://maxroll.gg/">maxroll.gg</a></td>
            <td><a href="https://www.ory.sh/case-studies/maxroll">Maxroll Case Study</a></td>
        </tr>
        <tr>
            <td>Zezam</td>
            <td align="center">
                <picture>
                    <source media="(prefers-color-scheme: dark)" srcset="https://raw.githubusercontent.com/ory/meta/master/static/adopters/light/zezam.svg" />
                    <img height="32px" src="https://raw.githubusercontent.com/ory/meta/master/static/adopters/dark/zezam.svg" alt="Zezam">
                </picture>
            </td>
            <td><a href="https://www.zezam.io/">zezam.io</a></td>
            <td><a href="https://www.ory.sh/case-studies/zezam">Zezam Case Study</a></td>
        </tr>
        <tr>
            <td>T.RowePrice</td>
            <td align="center">
                <picture>
                    <source media="(prefers-color-scheme: dark)" srcset="https://raw.githubusercontent.com/ory/meta/master/static/adopters/light/troweprice.svg" />
                    <img height="32px" src="https://raw.githubusercontent.com/ory/meta/master/static/adopters/dark/troweprice.svg" alt="T.RowePrice">
                </picture>
            </td>
            <td><a href="https://www.troweprice.com/">troweprice.com</a></td>
        </tr>
        <tr>
            <td>Mistral</td>
            <td align="center">
                <picture>
                    <source media="(prefers-color-scheme: dark)" srcset="https://raw.githubusercontent.com/ory/meta/master/static/adopters/light/mistral.svg" />
                    <img height="32px" src="https://raw.githubusercontent.com/ory/meta/master/static/adopters/dark/mistral.svg" alt="Mistral">
                </picture>
            </td>
            <td><a href="https://www.mistral.ai/">mistral.ai</a></td>
        </tr>
        <tr>
            <td>Axel Springer</td>
            <td align="center">
                <picture>
                    <source media="(prefers-color-scheme: dark)" srcset="https://raw.githubusercontent.com/ory/meta/master/static/adopters/light/axelspringer.svg" />
                    <img height="22px" src="https://raw.githubusercontent.com/ory/meta/master/static/adopters/dark/axelspringer.svg" alt="Axel Springer">
                </picture>
            </td>
            <td><a href="https://www.axelspringer.com/">axelspringer.com</a></td>
        </tr>
        <tr>
            <td>Hemnet</td>
            <td align="center">
                <picture>
                    <source media="(prefers-color-scheme: dark)" srcset="https://raw.githubusercontent.com/ory/meta/master/static/adopters/light/hemnet.svg" />
                    <img height="32px" src="https://raw.githubusercontent.com/ory/meta/master/static/adopters/dark/hemnet.svg" alt="Hemnet">
                </picture>
            </td>
            <td><a href="https://www.hemnet.se/">hemnet.se</a></td>
        </tr>
        <tr>
            <td>Cisco</td>
            <td align="center">
                <picture>
                    <source media="(prefers-color-scheme: dark)" srcset="https://raw.githubusercontent.com/ory/meta/master/static/adopters/light/cisco.svg" />
                    <img height="32px" src="https://raw.githubusercontent.com/ory/meta/master/static/adopters/dark/cisco.svg" alt="Cisco">
                </picture>
            </td>
            <td><a href="https://www.cisco.com/">cisco.com</a></td>
        </tr>
        <tr>
            <td>Presidencia de la República Dominicana</td>
            <td align="center">
                <picture>
                    <source media="(prefers-color-scheme: dark)" srcset="https://raw.githubusercontent.com/ory/meta/master/static/adopters/light/republica-dominicana.svg" />
                    <img height="42px" src="https://raw.githubusercontent.com/ory/meta/master/static/adopters/dark/republica-dominicana.svg" alt="Presidencia de la República Dominicana">
                </picture>
            </td>
            <td><a href="https://www.presidencia.gob.do/">presidencia.gob.do</a></td>
        </tr>
        <tr>
            <td>Moonpig</td>
            <td align="center">
                <picture>
                    <source media="(prefers-color-scheme: dark)" srcset="https://raw.githubusercontent.com/ory/meta/master/static/adopters/light/moonpig.svg" />
                    <img height="32px" src="https://raw.githubusercontent.com/ory/meta/master/static/adopters/dark/moonpig.svg" alt="Moonpig">
                </picture>
            </td>
            <td><a href="https://www.moonpig.com/">moonpig.com</a></td>
        </tr>
        <tr>
            <td>Booster</td>
            <td align="center">
                <picture>
                    <source media="(prefers-color-scheme: dark)" srcset="https://raw.githubusercontent.com/ory/meta/master/static/adopters/light/booster.svg" />
                    <img height="18px" src="https://raw.githubusercontent.com/ory/meta/master/static/adopters/dark/booster.svg" alt="Booster">
                </picture>
            </td>
            <td><a href="https://www.choosebooster.com/">choosebooster.com</a></td>
        </tr>
        <tr>
            <td>Zaptec</td>
            <td align="center">
                <picture>
                    <source media="(prefers-color-scheme: dark)" srcset="https://raw.githubusercontent.com/ory/meta/master/static/adopters/light/zaptec.svg" />
                    <img height="24px" src="https://raw.githubusercontent.com/ory/meta/master/static/adopters/dark/zaptec.svg" alt="Zaptec">
                </picture>
            </td>
            <td><a href="https://www.zaptec.com/">zaptec.com</a></td>
        </tr>
        <tr>
            <td>Klarna</td>
            <td align="center">
                <picture>
                    <source media="(prefers-color-scheme: dark)" srcset="https://raw.githubusercontent.com/ory/meta/master/static/adopters/light/klarna.svg" />
                    <img height="24px" src="https://raw.githubusercontent.com/ory/meta/master/static/adopters/dark/klarna.svg" alt="Klarna">
                </picture>
            </td>
            <td><a href="https://www.klarna.com/">klarna.com</a></td>
        </tr>
        <tr>
            <td>Raspberry PI Foundation</td>
            <td align="center">
                <picture>
                    <source media="(prefers-color-scheme: dark)" srcset="https://raw.githubusercontent.com/ory/meta/master/static/adopters/light/raspi.svg" />
                    <img height="32px" src="https://raw.githubusercontent.com/ory/meta/master/static/adopters/dark/raspi.svg" alt="Raspberry PI Foundation">
                </picture>
            </td>
            <td><a href="https://www.raspberrypi.org/">raspberrypi.org</a></td>
        </tr>
        <tr>
            <td>Tulip</td>
            <td align="center">
                <picture>
                    <source media="(prefers-color-scheme: dark)" srcset="https://raw.githubusercontent.com/ory/meta/master/static/adopters/light/tulip.svg" />
                    <img height="32px" src="https://raw.githubusercontent.com/ory/meta/master/static/adopters/dark/tulip.svg" alt="Tulip Retail">
                </picture>
            </td>
            <td><a href="https://tulip.com/">tulip.com</a></td>
        </tr>
        <tr>
            <td>Hootsuite</td>
            <td align="center">
                <picture>
                    <source media="(prefers-color-scheme: dark)" srcset="https://raw.githubusercontent.com/ory/meta/master/static/adopters/light/hootsuite.svg" />
                    <img height="32px" src="https://raw.githubusercontent.com/ory/meta/master/static/adopters/dark/hootsuite.svg" alt="Hootsuite">
                </picture>
            </td>
            <td><a href="https://hootsuite.com/">hootsuite.com</a></td>
        </tr>
        <tr>
            <td>Segment</td>
            <td align="center">
                <picture>
                    <source media="(prefers-color-scheme: dark)" srcset="https://raw.githubusercontent.com/ory/meta/master/static/adopters/light/segment.svg" />
                    <img height="32px" src="https://raw.githubusercontent.com/ory/meta/master/static/adopters/dark/segment.svg" alt="Segment">
                </picture>
            </td>
            <td><a href="https://segment.com/">segment.com</a></td>
        </tr>
        <tr>
            <td>Arduino</td>
            <td align="center">
                <picture>
                    <source media="(prefers-color-scheme: dark)" srcset="https://raw.githubusercontent.com/ory/meta/master/static/adopters/light/arduino.svg" />
                    <img height="32px" src="https://raw.githubusercontent.com/ory/meta/master/static/adopters/dark/arduino.svg" alt="Arduino">
                </picture>
            </td>
            <td><a href="https://www.arduino.cc/">arduino.cc</a></td>
        </tr>
        <tr>
            <td>Sainsbury's</td>
            <td align="center">
                <picture>
                    <source media="(prefers-color-scheme: dark)" srcset="https://raw.githubusercontent.com/ory/meta/master/static/adopters/light/sainsburys.svg" />
                    <img height="24px" src="https://raw.githubusercontent.com/ory/meta/master/static/adopters/dark/sainsburys.svg" alt="Sainsbury's">
                </picture>
            </td>
            <td><a href="https://www.sainsburys.co.uk/">sainsburys.co.uk</a></td>
        </tr>
        <tr>
            <td>Contraste</td>
            <td align="center">
                <picture>
                    <source media="(prefers-color-scheme: dark)" srcset="https://raw.githubusercontent.com/ory/meta/master/static/adopters/light/contraste.svg" />
                    <img height="32px" src="https://raw.githubusercontent.com/ory/meta/master/static/adopters/dark/contraste.svg" alt="Contraste">
                </picture>
            </td>
            <td><a href="https://www.contraste.com/en">contraste.com</a></td>
        </tr>
        <tr>
            <td>inMusic</td>
            <td align="center">
                <picture>
                    <source media="(prefers-color-scheme: dark)" srcset="https://raw.githubusercontent.com/ory/meta/master/static/adopters/light/inmusic.svg" />
                    <img height="24px" src="https://raw.githubusercontent.com/ory/meta/master/static/adopters/dark/inmusic.svg" alt="InMusic">
                </picture>
            </td>
            <td><a href="https://inmusicbrands.com/">inmusicbrands.com</a></td>
        </tr>
        <tr>
            <td>Buhta</td>
            <td align="center">
                <picture>
                    <source media="(prefers-color-scheme: dark)" srcset="https://raw.githubusercontent.com/ory/meta/master/static/adopters/light/buhta.svg" />
                    <img height="32px" src="https://raw.githubusercontent.com/ory/meta/master/static/adopters/dark/buhta.svg" alt="Buhta">
                </picture>
            </td>
            <td><a href="https://buhta.com/">buhta.com</a></td>
        </tr>
        </tr>
            <tr>
            <td>Amplitude</td>
            <td align="center">
                <picture>
                    <source media="(prefers-color-scheme: dark)" srcset="https://raw.githubusercontent.com/ory/meta/master/static/adopters/light/amplitude.svg" />
                    <img height="32px" src="https://raw.githubusercontent.com/ory/meta/master/static/adopters/dark/amplitude.svg" alt="amplitude.com">
                </picture>
            </td>
            <td><a href="https://amplitude.com/">amplitude.com</a></td>
        </tr>
    <tr>
      <td align="center"><a href="https://tier4.jp/en/"><picture><source media="(prefers-color-scheme: dark)" srcset="https://raw.githubusercontent.com/ory/meta/master/static/adopters/light/tieriv.svg" /><img height="32px" src="https://raw.githubusercontent.com/ory/meta/master/static/adopters/dark/tieriv.svg" alt="TIER IV"></picture></a></td>
      <td align="center"><a href="https://kyma-project.io"><picture><source media="(prefers-color-scheme: dark)" srcset="https://raw.githubusercontent.com/ory/meta/master/static/adopters/light/kyma.svg" /><img height="32px" src="https://raw.githubusercontent.com/ory/meta/master/static/adopters/dark/kyma.svg" alt="Kyma Project"></picture></a></td>
      <td align="center"><a href="https://serlo.org/"><picture><source media="(prefers-color-scheme: dark)" srcset="https://raw.githubusercontent.com/ory/meta/master/static/adopters/light/serlo.svg" /><img height="32px" src="https://raw.githubusercontent.com/ory/meta/master/static/adopters/dark/serlo.svg" alt="Serlo"></picture></a></td>
      <td align="center"><a href="https://padis.io/"><picture><source media="(prefers-color-scheme: dark)" srcset="https://raw.githubusercontent.com/ory/meta/master/static/adopters/light/padis.svg" /><img height="32px" src="https://raw.githubusercontent.com/ory/meta/master/static/adopters/dark/padis.svg" alt="Padis"></picture></a></td>
    </tr>
    <tr>
      <td align="center"><a href="https://cloudbear.eu/"><picture><source media="(prefers-color-scheme: dark)" srcset="https://raw.githubusercontent.com/ory/meta/master/static/adopters/light/cloudbear.svg" /><img height="32px" src="https://raw.githubusercontent.com/ory/meta/master/static/adopters/dark/cloudbear.svg" alt="Cloudbear"></picture></a></td>
      <td align="center"><a href="https://securityonionsolutions.com/"><picture><source media="(prefers-color-scheme: dark)" srcset="https://raw.githubusercontent.com/ory/meta/master/static/adopters/light/securityonion.svg" /><img height="32px" src="https://raw.githubusercontent.com/ory/meta/master/static/adopters/dark/securityonion.svg" alt="Security Onion Solutions"></picture></a></td>
      <td align="center"><a href="https://factlylabs.com/"><picture><source media="(prefers-color-scheme: dark)" srcset="https://raw.githubusercontent.com/ory/meta/master/static/adopters/light/factly.svg" /><img height="24px" src="https://raw.githubusercontent.com/ory/meta/master/static/adopters/dark/factly.svg" alt="Factly"></picture></a></td>
      <td align="center"><a href="https://cashdeck.com.au/"><picture><source media="(prefers-color-scheme: dark)" srcset="https://raw.githubusercontent.com/ory/meta/master/static/adopters/light/allmyfunds.svg" /><img height="32px" src="https://raw.githubusercontent.com/ory/meta/master/static/adopters/dark/allmyfunds.svg" alt="All My Funds"></picture></a></td>
    </tr>
    <tr>
      <td align="center"><a href="https://nortal.com/"><picture><source media="(prefers-color-scheme: dark)" srcset="https://raw.githubusercontent.com/ory/meta/master/static/adopters/light/nortal.svg" /><img height="32px" src="https://raw.githubusercontent.com/ory/meta/master/static/adopters/dark/nortal.svg" alt="Nortal"></picture></a></td>
      <td align="center"><a href="https://www.ordermygear.com/"><picture><source media="(prefers-color-scheme: dark)" srcset="https://raw.githubusercontent.com/ory/meta/master/static/adopters/light/ordermygear.svg" /><img height="32px" src="https://raw.githubusercontent.com/ory/meta/master/static/adopters/dark/ordermygear.svg" alt="OrderMyGear"></picture></a></td>
      <td align="center"><a href="https://r2devops.io/"><picture><source media="(prefers-color-scheme: dark)" srcset="https://raw.githubusercontent.com/ory/meta/master/static/adopters/light/r2devops.svg" /><img height="32px" src="https://raw.githubusercontent.com/ory/meta/master/static/adopters/dark/r2devops.svg" alt="R2Devops"></picture></a></td>
      <td align="center"><a href="https://www.paralus.io/"><picture><source media="(prefers-color-scheme: dark)" srcset="https://raw.githubusercontent.com/ory/meta/master/static/adopters/light/paralus.svg" /><img height="32px" src="https://raw.githubusercontent.com/ory/meta/master/static/adopters/dark/paralus.svg" alt="Paralus"></picture></a></td>
    </tr>
    <tr>
      <td align="center"><a href="https://dyrector.io/"><picture><source media="(prefers-color-scheme: dark)" srcset="https://raw.githubusercontent.com/ory/meta/master/static/adopters/light/dyrector_io.svg" /><img height="32px" src="https://raw.githubusercontent.com/ory/meta/master/static/adopters/dark/dyrector_io.svg" alt="dyrector.io"></picture></a></td>
      <td align="center"><a href="https://pinniped.dev/"><picture><source media="(prefers-color-scheme: dark)" srcset="https://raw.githubusercontent.com/ory/meta/master/static/adopters/light/pinniped.svg" /><img height="32px" src="https://raw.githubusercontent.com/ory/meta/master/static/adopters/dark/pinniped.svg" alt="pinniped.dev"></picture></a></td>
      <td align="center"><a href="https://pvotal.tech/"><picture><source media="(prefers-color-scheme: dark)" srcset="https://raw.githubusercontent.com/ory/meta/master/static/adopters/light/pvotal.svg" /><img height="32px" src="https://raw.githubusercontent.com/ory/meta/master/static/adopters/dark/pvotal.svg" alt="pvotal.tech"></picture></a></td>
      <td></td>
    </tr>
    </tbody>
</table>

Many thanks to all individual contributors

<a href="https://opencollective.com/ory" target="_blank"><img src="https://opencollective.com/ory/contributors.svg?width=890&limit=714&button=false" /></a>

<!--END ADOPTERS-->

## Installation

Head over to the documentation to learn about ways of
[installing Ory Keto](https://www.ory.sh/docs/keto/install).

## Ecosystem

<!--BEGIN ECOSYSTEM-->

We build Ory on several guiding principles when it comes to our architecture
design:

- Minimal dependencies
- Runs everywhere
- Scales without effort
- Minimize room for human and network errors

Ory's architecture is designed to run best on a Container Orchestration system
such as Kubernetes, CloudFoundry, OpenShift, and similar projects. Binaries are
small (5-15MB) and available for all popular processor types (ARM, AMD64, i386)
and operating systems (FreeBSD, Linux, macOS, Windows) without system
dependencies (Java, Node, Ruby, libxml, ...).

### Ory Kratos: Identity and User Infrastructure and Management

[Ory Kratos](https://github.com/ory/kratos) is an API-first Identity and User
Management system that is built according to
[cloud architecture best practices](https://www.ory.sh/docs/next/ecosystem/software-architecture-philosophy).
It implements core use cases that almost every software application needs to
deal with: Self-service Login and Registration, Multi-Factor Authentication
(MFA/2FA), Account Recovery and Verification, Profile, and Account Management.

### Ory Hydra: OAuth2 & OpenID Connect Server

[Ory Hydra](https://github.com/ory/hydra) is an OpenID Certified™ OAuth2 and
OpenID Connect Provider which easily connects to any existing identity system by
writing a tiny "bridge" application. It gives absolute control over the user
interface and user experience flows.

### Ory Oathkeeper: Identity & Access Proxy

[Ory Oathkeeper](https://github.com/ory/oathkeeper) is a BeyondCorp/Zero Trust
Identity & Access Proxy (IAP) with configurable authentication, authorization,
and request mutation rules for your web services: Authenticate JWT, Access
Tokens, API Keys, mTLS; Check if the contained subject is allowed to perform the
request; Encode resulting content into custom headers (`X-User-ID`), JSON Web
Tokens and more!

### Ory Keto: Access Control Policies as a Server

[Ory Keto](https://github.com/ory/keto) is a policy decision point. It uses a
set of access control policies, similar to AWS IAM Policies, in order to
determine whether a subject (user, application, service, car, ...) is authorized
to perform a certain action on a resource.

<!--END ECOSYSTEM-->

## Security

### Disclosing Vulnerabilities

If you think you found a security vulnerability, please refrain from posting it
publicly on the forums, the chat, or GitHub. You can find all info for
responsible disclosure in our
[security.txt](https://www.ory.sh/.well-known/security.txt).

## Telemetry

Our services collect summarized, anonymized data which can optionally be turned
off. Click [here](https://www.ory.sh/docs/ecosystem/sqa) to learn more.

### Guide

The Guide is available [here](https://www.ory.sh/docs/keto/).

### HTTP API Documentation

The HTTP API is documented [here](https://www.ory.sh/docs/keto/sdk/api).

### Upgrading and Changelog

New releases might introduce breaking changes. To help you identify and
incorporate those changes, we document these changes in
[UPGRADE.md](./UPGRADE.md) and [CHANGELOG.md](./CHANGELOG.md).

### Command Line Documentation

Run `keto -h` or `keto help`.

### Develop

We encourage all contributions and recommend you read our
[contribution guidelines](./CONTRIBUTING.md).

#### Dependencies

You need Go 1.19+ and (for the test suites):

- Docker and Docker Compose
- GNU Make 4.3
- NodeJS / npm >= v7

It is possible to develop Ory Keto on Windows, but please be aware that all
guides assume a Unix shell like bash or zsh.

#### Install From Source

<pre type="make/command">
make install
</pre>

#### Formatting Code

You can format all code using <code type="make/command">make format</code>. Our
CI checks if your code is properly formatted.

#### Running Tests

There are two types of tests you can run:

- Short tests (do not require a SQL database like PostgreSQL)
- Regular tests (do require PostgreSQL, MySQL, CockroachDB)

##### Short Tests

Short tests run fairly quickly. You can either test all of the code at once:

```shell script
go test -short -tags sqlite ./...
```

or test just a specific module:

```shell script
go test -tags sqlite -short ./internal/check/...
```

##### Regular Tests

Regular tests require a database set up. Our test suite is able to work with
docker directly (using [ory/dockertest](https://github.com/ory/dockertest)) but
we encourage to use the script instead. Using dockertest can bloat the number of
Docker Images on your system and starting them on each run is quite slow.
Instead we recommend doing:

```shell
source ./scripts/test-resetdb.sh
go test -tags sqlite ./...
```

##### End-to-End Tests

The e2e tests are part of the normal `go test`. To only run the e2e test, use:

```shell
source ./scripts/test-resetdb.sh
go test -tags sqlite ./internal/e2e/...
```

or add the `-short` tag to only test against sqlite in-memory.

#### Build Docker

You can build a development Docker Image using:

<pre type="make/command">
make docker
</pre>
