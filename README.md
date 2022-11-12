<h1 align="center"> Ticketing Project </h1> <br>
<p align="center">
  <a href="https://gitpoint.co/">
    <img alt="GitPoint" title="GitPoint" src="http://i.imgur.com/VShxJHs.png" width="450">
  </a>
</p>

<p align="center">
  This project is a project that provides a simple ticket reservation service on the website.
</p>


[![Build Status](https://img.shields.io/travis/gitpoint/git-point.svg?style=flat-square)](https://travis-ci.org/gitpoint/git-point)
[![Coveralls](https://img.shields.io/coveralls/github/gitpoint/git-point.svg?style=flat-square)](https://coveralls.io/github/gitpoint/git-point)
[![All Contributors](https://img.shields.io/badge/all_contributors-73-orange.svg?style=flat-square)](./CONTRIBUTORS.md)
[![PRs Welcome](https://img.shields.io/badge/PRs-welcome-brightgreen.svg?style=flat-square)](http://makeapullrequest.com)
[![Commitizen friendly](https://img.shields.io/badge/commitizen-friendly-brightgreen.svg?style=flat-square)](http://commitizen.github.io/cz-cli/)
[![Gitter chat](https://img.shields.io/badge/chat-on_gitter-008080.svg?style=flat-square)](https://gitter.im/git-point)


<!-- START doctoc generated TOC please keep comment here to allow auto update -->
<!-- DON'T EDIT THIS SECTION, INSTEAD RE-RUN doctoc TO UPDATE -->
## Table of Contents

- [사이드 프로젝트 진행하는 목적](#Goals)
- [프로젝트 기능들](#Features)
- [AWS 인프라 설계도](#Infrastructure Design)
- [Getting started](#Getting started)
- [Installation](#Installation)
- [Configuration](#Configuration)
- [Environment variables](#Environment variables)


<!-- END doctoc generated TOC please keep comment here to allow auto update -->


## Goals
* 인프라 설계 및 구축 (무중단 배포)
* MySQL vs MongoDB (ORM 사용을 통한 차이점 느끼기)
* Unit Test 작성
* CI/CD 구축
* 클린 아키텍처 적용
* 공통 모듈 제대로 사용해보기

## Features

These are the back-end work done on the project:

* 티겟 예매하기
* OAuth 로그인
* Import 결제하기
* 배치 시스템 (결제 전 알림)


## Infrastructure Design

The figure below shows the entire infrastructure.
## Backend Infrastructure.

<p align="center">
    <img width="1307" alt="스크린샷 2022-11-12 오전 3 41 26" src="https://user-images.githubusercontent.com/85932211/201409200-a7a69173-e255-4150-a148-9ee41ec6890c.png">
</p>

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
