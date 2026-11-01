# club soccer schedule API

Kids club soccer made easy.

A Go service that parses GotSport team schedule pages, extracts the games, and serves them as JSON from a small HTTP API.

## How it works

```
GotSport schedule HTML  ->  parser (goquery)  ->  []Game  ->  JSON  ->  /schedule endpoint
```

The parser reads schedule pages saved as `.html` files. GotSport blocks automated requests with a CAPTCHA and its `robots.txt` disallows bots, so the project does not fetch pages itself. You save a schedule from your browser into `data/`, and the service does the rest. The parser is independent of where the HTML comes from, so a sanctioned data source (a feed or permission from GotSport) would only replace the input step.

## Parsing

On a team schedule page (`/org_event/events/<eventID>/schedules?team=<teamID>`):

- Each game is a `<tr class='fz-sm'>` in the desktop table. The mobile layout repeats every game, so selecting `tr.fz-sm` avoids duplicates, and the standings table (which shares CSS classes) is never matched.
- Each row has 7 cells:

| Cell | Contents |
|---|---|
| 0 | Match number |
| 1 | Date, time, and status |
| 2 | Home team |
| 3 | Result (`2 - 2`, or `-` if not played) |
| 4 | Away team |
| 5 | Location |
| 6 | Division |

## Run

```sh
go mod download
go run .
```

## Status

- [x] Load saved HTML into goquery
- [ ] Parse game rows into a `Game` struct
- [ ] JSON output
- [ ] Filter by team and upcoming weekend
- [ ] Tests against saved fixtures
- [ ] `/schedule` HTTP endpoint
- [ ] Load every `.html` file in `data/` to support multiple clubs

## Files

- `main.go` - entry point
- `game.go` - the `Game` type
- `schedule.html` - sample saved schedule page
