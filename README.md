![fullsizebanner](https://user-images.githubusercontent.com/96031819/233515645-147df882-0f6e-427c-89e5-a13e3c053a4a.png)

![donuts-are-good's followers](https://img.shields.io/github/followers/donuts-are-good?&color=555&style=for-the-badge&label=followers) ![donuts-are-good's stars](https://img.shields.io/github/stars/donuts-are-good?affiliations=OWNER%2CCOLLABORATOR&color=555&style=for-the-badge) ![donuts-are-good's visitors](https://komarev.com/ghpvc/?username=donuts-are-good&color=555555&style=for-the-badge&label=visitors)

# webapp login authentication boilerplate

this is a simple login authentication boilerplate written in go that can be used as a starting point for your webapp.

## features

- written in stdlib go
- secure authentication with bcrypt password hashing
- csrf protection using cookies and tokens
- basic email validation
- simple html templates for login, registration, and 404 error pages


## getting started

clone the code:
```bash
git clone https://github.com/donuts-are-good/login-auth-boilerplate.git
```

enter the cloned folder:
```bash
cd login-auth-boilerplate
```

build & run:
```shell
go build && ./login-auth-boilerplate
```

open a web browser and go to:

-  `http://localhost:8080` to see the **main** page.
-  `http://localhost:8080/signup` to see the **signup** page.
-  `http://localhost:8080/login` to see the **login** page.

## usage

you can use this boilerplate as the first building blocks of your program. i made it because i often repeat these patterns, to varying degrees of correctness and just wanted to do my best one last time and then use that as the basis for other apps every time after, to save effort.

the code creates several endpoints:

-  **/ -**  the index page.
-  **/login -**  the login page. accepts post requests to authenticate users.
-  **/register -**  the registration page. accepts post requests to create new user accounts.

**404.html**

the code also has a 404 page. the way the 404 page works is that any time a url that we don't understand is encountered, we send it to the index page. the index page handler then looks at the request path that brought you to be looking at it, and if it doesn't match `/` it will send you to the 404 handler, which corresponds to `templates/404.html`.

**index.html**

the index page is your landing page. generally i'll put some marketing content there and then make a nav. i don't know how other people want their apps to look so i stripped out anything that didn't need to be there that you might not want. 

**register.html**

this is just your basic email and password signup form. you'll notice there's no "remember me" or second password verify box. i go back and forth on whether or not to verify the password these days, or to [even have signups](https://github.com/donuts-are-good/signupless) [at all](https://github.com/donuts-are-good/libsignupless), however i wanted to leave that up to the end user.

**login.html**

this login form should be modified in the backend so that when the user succeeds in loggin in that they get sent to another route, whereas currently we just print a message welcoming the user.


## things you might not notice it does
here's a short list of things that this boilerplate takes into account that are good things you might not notice at first when running it.

- **email validation.** people put all kinds of crazy stuff in email fields. this is the stdlib way to check for that.
- **using bcrypt.** some people store passwords in plaintext, we don't do that
- **using constant time comparison.** i've got to be honest, nobody's likely to try and attack you this way, but that's no longer an issue. this is to resist against timing attacks.
- **using a sha256 prehash.** when we get a password, we pre-hash it with sha256 to be able to send a dependable length value to bcrypt, due to bcrypt having a 72byte limit on input.
- **csrf tokens.** these are to prevent cross-site request forgery (csrf). we do that by generating a unique token in a cookie for each form and verifying it when the form is submitted. 
- **http-only cookies.**  the csrf token cookie is set to http-only to prevent client-side scripts from scanning your token down


## changes you (might) want to make
this is just a boilerplate, and while it does more than what a lot of web apps might be doing, it is far from complete for some purposes. here is a brief list of things you might want to google about later when using this boilerplate:

- **using a database.** right now, if you relaunch the app, you lose all your signups.
- **forgot-password.** there is no way to handle password resets right now in this code
- **remember me.** we're storing cookies, so we could make it optional if you wanted to add this. 
- **cookie middleware.** this seemed outside of the scope of this program, but ideally you want to preface any request to a privileged user-bound area with a check for the correct cookie, but it seemed like to do this the way i'd normally do it i'd be implementing too much that doesn't belong in a simple boilerplate like this. 
- **password complexity.** generall i try not to involve javascript, and though this can be handled in the backend, i go back and forth on whether or not forcing a certain password schema makes it easier or harder to brute force a login.
- **rate limiting.** if someone wanted to brute force your login through this service, they will.
- **inpput field hygeine.** we're not checking the input fields for excess length or other less-common forms of abuse.


## license
mit license donuts-are-good https://github.com/donuts-are-good

see [license.md](license.md) if you want to know more. if you don't know what it means, don't sweat it. enjoy the code :)
