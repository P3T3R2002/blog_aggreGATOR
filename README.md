# blog_aggreGATOR

You will need Postgres and Go installed.

To install, copy ```go install github.com/P3T3R2002/blog_aggreGATOR@latest``` to the command line.

At first use it will create a .gatorconfig.json at the home directory, if it does not already exists, with the dbURL inside.

Command list:
* reset:        Resets the database.
* register:     Register a user.
* login:        Change user.
* users:        Print out all registered users.
* addfeed:      Creates a new feed. Needs a feed name and an url.
* feeds:        Print out all created feeds.
* follow:       Follow a feed. Need url.
* unfollow:     Unfollow a feed. Need url.
* following:    Print out all followed feed.
* agg:          Looks over every feed created and creates new Posts. Need a delay as duration string(1s, 1m, 1h).
* browse:       Lists new posts created by the agg command. Needs a number as the number of Posts to list. Default is 2.