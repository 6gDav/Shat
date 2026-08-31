# Shat

<p align="center">
    <img src="components/docs/appIcon.png" alt="Shat Icon" width="200">
</p>

# ❓ What is this

This is a simple light-weight single-use [LAN](https://en.wikipedia.org/wiki/Local_area_network) based chat application.

Click here 👉 [to navigate to user features](#forusers) <br>
Click here 👉 [to navigate to developer features](#fordevelopers)

# 💻 About this project

This application focuses on fixing call lag caused by slow internet connections. Take school classes as an example: during coding lessons, teachers often share their screens via [Microsoft Teams](https://www.microsoft.com/en-us/microsoft-teams/log-in), but due to poor connection speeds, the stream lags or freezes completely.

> ❗️ The last feature is still missing because I've been focusing on text transfer rather than video calling, but I plan to implement the video call function as well. 

## <a name="forusers"></a>⚡️ Core Features

This is a chat application that contains all the necessary features to work properly, including:
* ༕ **Registration** - Easy registration, only your name is needed.
* ✉️ **Private Messaging** - Real-time messaging with your friends.
* 📭 **Group Messaging** - Real-time messaging with all your mates.

## 📦 How to access

The application itself is only available in a desktop version. <br>
Navigate to the Releases menu on the left and choose the appropriate package.

## 🤔 How to use

### 🛜 Hoster 

This application provides a star-topology client-server connection. <br>
This means that one of the nodes acts as the server.

<p align="center">
    <img src="components/docs/star.png" alt="Star-topology" width="200">
</p>

<hr>

After opening the app this screen will greet you.

<p align="center">
    <img src="components/docs/greedingscreen.png" alt="Star-topology" width="200">
</p>

There is 2 option what you see here.

1. **Start Button** – Starts the server.
   * Once the server starts, this button will show *Stop*.
   * Clicking the *Stop* button will shut down the server.

2. **Generate QR Code** - This button will generate a readable QR code in a niew window.

3. **Log Message Container** - Logs, notifications, and errors will be shown here.
    * The website link will also be shown here.

<hr>

### 🧑‍💻 User

After navigation to the hosted page by: <br>
- **Scanning the QR code**  <br>
Or
- **Navigating to the [`http://shat.local:3000`](http://shat.local:3000) in a browser:** 

This page will appear in front of you.

<p align="center">
    <img src="components/docs/regpage.png" alt="Registration page" width="200">
</p>

In the input field, just enter an ID or something that is specific to you. 

<hr>

By pressing Enter, you will be navigated to this page.

<p align="center">
    <img src="components/docs/dashboard.png" alt="Dashboard page" width="200">
</p>

Here you find a possibility to cahnge your name. <br>

<p align="center">
    <img src="components/docs/changeuser.png" alt="Change user name" width="200">
</p>

You will see here the *group* chat and a list of the other users.

<p align="center">
    <img src="components/docs/userlist.png" alt="User list" width="200">
</p>

Select the targeted user you want to chat with.  <br>
Or choose the *group* chat if you want to chat with every one. <br>
> There is a possibility to talk with yourself.

<hr>

Once you select a user, this will appear in front of you.

<p align="center">
    <img src="components/docs/chat.png" alt="Chat" width="200">
</p>

This layout is very similar to the other one, so you won't get lost.

<hr>

### 🤏 Small feauteres

* If you disconnect by closing the browser or just the tab, your progress won't be lost. However, your name will disappear from the user list, and your friends won't be able to send you messages while you are away.

* If you return, your data will remain intact.

* You cannot re-register unless the host stops the server, which will reset all processes.

## <a name="fordevelopers"></a> 🛠️ Developer Guidance

**Just a heads up:** *like 99% of developers out there, I used AI and premade code during development.*

## 📚 Tech Stack 

* **[Fyne](https://fyne.io/)** (The main application and its UI.)
    * **[Go](https://go.dev/)** (The application and the backend programming language.)
* **[SvelteKit](https://svelte.dev/docs/kit/introduction)** (Framework of the web page.)
   * **[TypeScript](https://www.typescriptlang.org/)** (Main programming language of the web page.)

*In the beginning, I wanted to build something lightweight that didn't consume a lot of resources, with easy-to-understand code, which is why I chose Go and Svelte over Rust and React.*