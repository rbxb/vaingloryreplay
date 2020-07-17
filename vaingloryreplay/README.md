# vaingloryreplay

This package makes it easier save and later play back Vainglory replays.  
Works on Windows.  
It *should* work similarly on MacOS.  

## Installing

1. You can compile the tools using Go or download the binaries here:  
  [windows_amd64.zip](windows_amd64.zip)  

2. Set the path for your tools.  
  If you compiled the tools yourself you should already have your Go/bin saved as a Path.  
  If you downloaded the tools, move the executable files (**vgrsave.exe** and **vgrplay.exe**) to where you want to keep them and set their location as a Path environment variable.  
  [How to set an environment variable](https://www.google.com/search?q=how+to+set+environment+variables)  

3. Find the *Temp* directory, which is where Vainglory replay files are stored.  
  On Windows, this directory located at `C:\Users\user\AppData\Local\Temp`.  

## Save a replay

Play a Vainglory match. When the match ends, wait on the results screen. **Do not exit the match yet.**  
Using the command line, make sure you are in the *Temp* directory, then run **vgrsave**.  
```shell
$ cd C:\Users\user\AppData\Local\Temp
$ vgrsave
```
You should see a new directory named *vainglory-replays* appear in the *Temp* directory.  
The *vainglory-replays* directory should be populated with the vgr files from your match.  
If this worked, you can exit the match.  

#### vgrsave flags

##### `-source`
The directory with the source vgr files.  
Defaults to the current directory.  

##### `-name`
The name of the replay to save.  
Defaults to the most recently modified replay in the *source* directory.  

##### `-save`
The directory where the replay will be saved.  
Defaults to `./vainglory-replays`.  

##### `-sname`
The name to save the replay as.  
Defaults to an auto generated name.  

## Play back a replay

You will need to initiate an in-game replay to be able to play back a previously saved replay. Start a solo practice match and surrender. Wait on the results screen.  
Using the command line, make sure you are in the *Temp* directory, then run **vgrplay**.  
```shell
$ cd C:\Users\user\AppData\Local\Temp
$ vgrplay
```
Click **replay** in-game and it should replay the saved match.  

#### vgrplay flags

##### `-source`
The directory with the source vgr files.  
Defaults to `./vainglory-replays`.  

##### `-sname`
The name of the replay to play.  
Defaults to the most recently modified replay in the *source* directory.  

##### `-overwrite`
The directory with the active vgr files.  
Defaults to the current directory.  

##### `-oname`
The name of the replay to overwrite.  
Defaults to the most recently modified replay in the *overwrite* directory.  
