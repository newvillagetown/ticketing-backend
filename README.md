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
- [클릭 아키텍처](#Clean_Architecture)
- [디렉토리 구조](#Directory_layout)
- [AWS 인프라 설계도](#Infrastructure_Design)
- [메시지큐 모델 디자인](#Message_Queue_Model_Design)
- [Getting started](#Getting_started)
- [Installation](#Installation)


<!-- END doctoc generated TOC please keep comment here to allow auto update -->


## Goals
* 인프라 설계 및 구축 (무중단 배포)
* MySQL vs MongoDB (ORM 사용을 통한 차이점 느끼기)
* Unit Test 작성
* CI/CD 구축
* 클린 아키텍처 적용
* 공통 모듈 제대로 사용해보기
* Go 컨벤션 작업을 통한 Go스럽게 사용해보기
* 데이터 스트리밍 파이프라인 구축 (Goroutine + channel)

## Clean_Architecture

아키텍처 적용시 고려사항
- 계층을 분리하여 관심사를 분리한다.
- 특정 소프트웨어 라이브러리에 의존하지 않는다.
- 비즈니스 규칙은 UI, 데이터베이스, 웹 서버 또는 기타 외부 요소 없이 테스트 할 수 있어야 한다.
- 비즈니스 로직은 외부 세계와 무관하게 작동할 수 있어야 한다.

This project has 4 Domain layer :
- Models Layer
- Repository Layer
- Usecase Layer
- Delivery Layer (Controller/Handler)

The diagram
<p align="center">
    <img width="990" alt="스크린샷 2022-12-22 오후 7 46 07" src="https://user-images.githubusercontent.com/35329247/209118510-3153c568-0d17-43de-a778-210dd53002c5.png">
</p>

## Directory_layout

    .
    ├── common                  # 서버 공통 모듈
    ├── deployment              # 배포 폴더 (Dockerfile, taskfile)
    ├── docs                    # swagger 폴더
    ├── features                # 기능 폴더
    │   ├── ....
    │   ├── product                     
    │   │   ├── domain                  # 도메인은 엔티티와 동일하며 모든 레이어에서 사용
    │   │   ├── handler                 # 프리젠테이션을 담당하는 영역으로 데이터가 표시되는 방식을 결정
    │   │   ├── repository              # 데이터를 관리
    │   │   └── usecase                 # 비즈니스 프로세스를 처리
    │   ├── ....
    └── ....

## Features

These are the back-end work done on the project:

* 티겟 예매하기
* OAuth 로그인
* Import 결제하기
* 배치 시스템 (결제 전 알림)


## Infrastructure_Design

The figure below shows the entire infrastructure.

<p align="center">
    <img width="1303" alt="스크린샷 2023-01-11 오전 2 19 39" src="https://user-images.githubusercontent.com/35329247/211793326-76e52dc4-9295-4ee0-b70a-dc553280b625.png">
</p>

## Message_Queue_Model_Design

비동기 이벤트 처리를 위한 메시지큐 디자인
<p align="center">
<img width="635" alt="message queue model" src="https://user-images.githubusercontent.com/85932211/209458485-dd7f3b62-291a-464e-9ef4-e91e3005bb05.png">
</p>

## Getting_started

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
