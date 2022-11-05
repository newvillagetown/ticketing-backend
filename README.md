<h1 align="center"> Ticketing Project </h1> <br>
<p align="center">
  <a href="https://gitpoint.co/">
    <img alt="GitPoint" title="GitPoint" src="http://i.imgur.com/VShxJHs.png" width="450">
  </a>
</p>

<p align="center">
  This project is a project that provides a simple ticket reservation service on the website.
</p>



<!-- START doctoc generated TOC please keep comment here to allow auto update -->
<!-- DON'T EDIT THIS SECTION, INSTEAD RE-RUN doctoc TO UPDATE -->
## Table of Contents

- [Infrastructure Design](#Infrastructure Design)
- [Features](#Features)
- [Getting started](#Getting started)
- [Installation](#Installation)
- [Configuration](#Configuration)
- [Environment variables](#Environment variables)


<!-- END doctoc generated TOC please keep comment here to allow auto update -->

## Infrastructure Design

[![Build Status](https://img.shields.io/travis/gitpoint/git-point.svg?style=flat-square)](https://travis-ci.org/gitpoint/git-point)
[![Coveralls](https://img.shields.io/coveralls/github/gitpoint/git-point.svg?style=flat-square)](https://coveralls.io/github/gitpoint/git-point)
[![All Contributors](https://img.shields.io/badge/all_contributors-73-orange.svg?style=flat-square)](./CONTRIBUTORS.md)
[![PRs Welcome](https://img.shields.io/badge/PRs-welcome-brightgreen.svg?style=flat-square)](http://makeapullrequest.com)
[![Commitizen friendly](https://img.shields.io/badge/commitizen-friendly-brightgreen.svg?style=flat-square)](http://commitizen.github.io/cz-cli/)
[![Gitter chat](https://img.shields.io/badge/chat-on_gitter-008080.svg?style=flat-square)](https://gitter.im/git-point)

The figure below shows the entire infrastructure.

**Backend Infrastructure.**

<p align="center">
  <img width="1370" alt="infra" src="https://user-images.githubusercontent.com/85932211/199739365-618e2f31-de19-4a28-9c1c-aa9813de113c.png">
</p>

## Features

These are the back-end work done on the project:

* Booking tickets
* Membership / membership withdrawal function
* Paying in conjunction with IMPORT
* Implementing an Alarm Placement System Using AWS Lambda


## Getting started

<pre>
$ air -c .air.toml
</pre>


## Installation

With go 1.16 or higher:
<pre>
$ go install github.com/cosmtrek/air@latest

# 1.16 or newer
$ go install github.com/swaggo/swag/cmd/swag@latest
</pre>
