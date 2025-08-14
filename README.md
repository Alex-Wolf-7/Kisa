# KISA #
Kisa is a League of Legends client application built to turn key bindings from per-account to per-champion.

### Why?: ###
I'm a Nami and Janna main.

Nami Q is an ability where the range indicator (quick cast with indicator) is very helpful.

Janna Q is an ability where the range indicator is a hinderance (it makes it difficult to double-tap Q and block dashes)

Swapping the "Q" key from QuickCast to QuickCast-With-Indicator and back every game is tedious, and leaves plenty of room for mistakes.

What if we could have Q always set to QuickCast-With-Indicator for Nami, and always set to QuickCast for Janna?

What about other uses? Do you like having range indicators for offensive spells but none for shields so they're quicker? Do you know a specific champion's ability ranges so well you don't want the indicators, but want them on other champions? Do you want to make sure you have selfcast available for heals and shields, but not when you're playing assassins?

This application is built to handle that all with as little disruption and effort as possible. It works for all champions, all cast types, and all 4 abilities. It currently does not function with Summoner Spells or Items, but if that's a desired feature let me know and it'd be easy to add.

### How to use: ###
Open the application. It will automatically detect and connect to your League of Legends client.

When your client is open, you can enter any champion name (case insensitive) to associate your client's current keybindings with that champion. Entering "default" will set your default keybindings, which are used for every champion that does not have specific keybindings set.

The better way to set champion-specific keybindings though is just to play games. At the end of every game, the application will save the keybindings you finished the game with, and associate them with the champion you just played.

Once keybindings are saved for a specific champion, whenver you start a new game with that champion those keybindings will be loaded. Saved keybindings will persist if you close and re-open the application.

That's it! Play games, set your keybindings to how you like them if they're ever off, and next time you play that champion the application will automatically update your client with the new keybindings.

### How to download: ###
Follow [this link](https://github.com/Alex-Wolf-7/Kisa/tree/meow/builds) to see all builds.

Builds are labeled by YYYY-MM-DD date. The newest builds should be at the bottom of this list.

If a build has the `.exe` extension, it is for Windows. If it does not have that extension, it is for Mac.

Builds with `_debug` have additional logging enabled to help track down issues. It performs the same as the non-debug version, but the application text output may fill more with information useless for normal use of this application.

Just download the newest version for your O.S. that doesn't have `_debug` (or does, if you want it), double click on it, and start a game!

### Feedback ###
This is a new application, and may have issues or bugs. If you encounter any issues, bugs, or have a good idea for some way to improve this application, please shoot me a message and let me know!

My Discord username is `CubicDolphin#6558`.

Happy League-ing!