# vainglorystream

This package extends [github.com/rbxb/vaingloryreplay](https://github.com/rbxb/vaingloryreplay) to make streaming Vainglory replays possible.
It will probably be helpful to read the vaingloryreplay README before reading this.

## Precompiled Binaries

[windows (amd64)](windows_amd64.zip)

## What it does

Vainglory has an in-game replay system that allows you to rewatch your matches. This package makes it possible to livestream a match in Vainglory's replay format so others can spectate the match using the in-game replay system. It's similar to Vainglory's spectator mode but you can have an unlimited number of viewers and you can spectate public matches.

## Using this package

There are three separate programs that work together to stream your replay:
- The server (vgsserver)
- The streaming client (vgsstream)
- The viewing client (vgsview)

## The server (vgsserver)
The server connects stream clients to view clients. You can have multiple streams hosted on the stream server (each stream will have a unique id to join).
Streamer clients must connect to `/stream` and viewer clients must connect to `/view`.

### Flags

#### `-port`
The port to listen on. (:8080)

## Streaming (vgsstream)
1. Start a Vainglory match.
2. Locate where the replay files are saved on your computer. On Windows, this folder is `C:\Users\user\AppData\Local\Temp`.
3. Run vgsstream. Use the address of the server and `/stream`
```shell
$ cd C:\Users\user\AppData\Local\Temp
$ vgsstream -address http://example.com/stream
```
4. You should receive a "Stream Code" from the server. Viewers will need to use this code to join your stream.

### Flags

#### `-source`
The directory with the source vgr files. (./)

#### `-name`
The name of the replay to save. (*picks the most recently modified replay*)

#### `-address`
The address of the server. (http://localhost:8080/stream)

## Viewing (vgsview)
1. Prepare a replay by starting a solo practice match and surrendering.
2. Locate where the replay files are saved on your computer. On Windows, this folder is `C:\Users\user\AppData\Local\Temp`.
3. Run vgsview. Use the Stream Code as the id.
```shell
$ cd C:\Users\user\AppData\Local\Temp
$ vgsview -address http://example.com/view -id 000000
```
4. Once a couple of replay frames have loaded, you can click **replay** in-game.

### Flags

#### `-address`
The address of the server. (http://localhost:8080/view)

#### `-id`
The stream id. (000000)

#### `-overwrite`
The directory with the active vgr files. (./)

#### `-oname`
The name of the replay to overwrite. (*picks the most recently modified replay*)
