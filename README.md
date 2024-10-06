

# Real-Time Forum

## Introduction

Real-Time Forum is a forum that allows users to communicate with each other in real time by creating posts, comments to them and personal messages. This project is a single-page application (SPA) uses the languages ​​Golang, JavaScript, HTML, CSS, as well as WebSockets and SQLite technologies

## Table of Contents
- [Installation](#installation)
- [Usage](#usage)
  - [Registration and Login](#registration-and-login)
  - [Creating Posts, Comments and messages](#creating-posts-comments-and-messages)
- [Configuration](#configuration)
- [Contributors](#contributors)


## Installation

- **Clone the repository:**

```bash
git clone https://github.com/Evgeny311/kood-real-time-forum
cd kood-real-time-forum
```

- **Set up the Backend (Golang):**

Ensure Go is installed and set up the Go workspace.

```bash
make run-api
```

- **Set up the Frontend(in separate terminal):**

```bash
make run-client
```

- **Database Setup:**

If there is no database, it is created and configured when the program is launched.

## Usage

### Registration and Login

- Users need to insert nickname, age, gender, first name, last name, email, and password for registration
- Login using  the nickname or email and the password.

### Creating Posts,Comments and messages

- Users can create posts and comment on them.
- Users can select categories for posts
- Posts are displayed in the feed
- Comments can be viewed by clicking on a post
- Users can use the chat to send private messages, display users online/offline, and reload past messages.
- Real-time messaging is done via WebSockets.

## Configuration

`config.json` is the storage for variables and configurations. You can change database paths, server ports and other parameters.

## Contributors

<div align="center">
  <table>
    <tbody><tr>
      <td align="center"><a href="https://01.kood.tech/git/eandreyc" rel="nofollow"></a></td>
      <td align="center"><a href="https://01.kood.tech/git/jkisselj" rel="nofollow"></a></td>
    </tr>
    <tr>
      <td align="center">eandreyc</td>
      <td align="center">jkisselj</td>
    </tr>
  </tbody></table>
</div>
