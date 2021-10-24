# Walpurgis 🔥

Command-line tool to query any social network and collect some stats and ease the effective clean-up removing inactive,
old and non-suitable accounts according your parameters.

## Configuration

```shell
> export WALPURGIS_PATH={APP_PATH}
> export API_KEY={TWITTER_TWITTER_API_KEY}
> export API_SECRET={TWITTER_API_SECRET}
> export ACCESS_KEY={TWITTER_ACCESS_KEY}
> export ACCESS_SECRET={TWITTER_ACCESS_SECRET}
```
Walpurgis uses a static configuration file. See example [here](./config/configuration.yml)
## Run

```shell
> go run .
```

## Commands

### List of available options: 
```shell
margostino@walpurgis> help

[ snapshot users ] - Collect and save an updated list of friends
[ help ] - List available commands
[ exit ] - Exit shell
[ show stats ] - Show general stats from user
[ rank users by x2 ] - Get friends ranking sorted by field(s)
[ select users where x3 ] - Select and filter users by condition(s)
```

### Take a snapshot of the users you follow
```shell
margostino@walpurgis> snapshot users
```

### Rank users by conditions (age, last activity, followers, following)
```shell
margostino@walpurgis> rank users by age asc

User: jack - Created At: 2006-03-21 20:50:14 +0000 +0000 - Last Activity: 2021-10-23 02:22:30 +0000 +0000]
User: sama - Created At: 2006-07-16 22:01:55 +0000 +0000 - Last Activity: 2021-10-22 15:48:07 +0000 +0000]
User: Werner - Created At: 2006-12-21 15:12:02 +0000 +0000 - Last Activity: 0001-01-01 00:00:00 +0000 UTC]
...
```

### Select users by conditions (text in description, email, status, name)
```shell
margostino@walpurgis> select users where description like climate

[jabeckx] - RT @aiyanabodi: American companies representing over 6.5M employees and 13% of the U.S. economy support strong climate investments—because…]
[UNDRR] - #UN Office for #Disaster #Risk Reduction %44%  dedicated to building disaster #resilience and tackling #climatechange through implementation of the Sendai Framework]
[GlobalEcoGuy] - RT @DavidRVetter: People talk a lot of crap about scientists not being able to communicate their work. I interview climate researchers day…]
[ClimateSignals] - Climate Signals provides resources in real time explaining how climate change worsens extreme weather and other impacts. A project of Climate Nexus.]
[AntarcticReport] - Follow for News on Antarctica and the Southern Ocean %44%  esp the hard science underlining the importance of Antarctica as a bellwether of global climate change.]
...
```
### Exit
```shell
margostino@walpurgis> exit
bye!
```

### TODO
- [ ] Validations (wrong or missing config)
- [ ] Preprocessing (users)
- [ ] Alias for recurrent commands
- [ ] Smart and more insightful queries
- [ ] Pattern matching
- [ ] Streaming processing


