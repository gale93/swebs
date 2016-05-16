![GitHub Logo](/media/logo.png)
A simple, ready to use, web server for tests ! 


## How to get it work ?

No dependencies, no nothing. Just click on the executable and you are ready.

#### I dont know how to click.

Here you go my little friend:

![How to Click](/media/swebs.gif)

## What it does ?

* Listen on a port
* Replay with the content of "resources" directory.
* Echo the JSON stuff you send via websocket ( /socket is the channel ) 

## Configs

You can change via flag this settings:

* -port ( Default: 9393 )
* -checkdir (Default: true. it creates the directory if not exists)
* -ws (Default: true. Enables the websocket thingy you read before )