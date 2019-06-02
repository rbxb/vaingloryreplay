# vaingloryreplay

This package makes it possible to save and later play back Vainglory replays.

## Precompiled Binaries

[windows (amd64)](windows_amd64.zip)

## Saving a replay (vgrsave)

Once the match is complete and **before you exit the match**, complete these steps:
1. Locate where the replay files are saved on your computer. On Windows, this folder is `C:\Users\user\AppData\Local\Temp`.
2. Run vgrsave
```shell
$ cd C:\Users\user\AppData\Local\Temp
$ vgrsave
```
3. You can now exit the match.

### Flags

#### `-source`
The directory with the source vgr files. (./)

#### `-name`
The name of the replay to save. (*picks the most recently modified replay*)

#### `-save`
The directory where the replay will be saved. (./vainglory-replays)

#### `-sname`
The name to save the replay as. (*auto generated*)

## Playing back a replay (vgrplay)

You will need to initiate an in-game replay to be able to play back a previously saved replay. Start a solo practice match and surrender. **Before clicking replay**, complete these steps:
1. Locate where the replay files are saved on your computer. On Windows, this folder is `C:\Users\user\AppData\Local\Temp`.
2. Locate where the saved replay is. If you didn't specify where to save the replay, it's probably in `Temp\vainglory-replays`.
3. Run vgrplay
```shell
$ cd C:\Users\user\AppData\Local\Temp
$ vgrplay
```
4. Click **replay** in-game and it should replay the saved match.

### Flags

#### `-source`
The directory with the source vgr files. (./vainglory-replays)

#### `-sname`
The name of the replay to play. (*picks the most recently modified replay*)

#### `-overwrite`
The directory with the active vgr files. (./)

#### `-oname`
The name of the replay to overwrite. (*picks the most recently modified replay*)
