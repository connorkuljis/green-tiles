# Green Tiles

🔨 Automatic contribution graph screenshot tool. Enter any github username and return an image of your current contribution graph.

## Demo

https://github.com/connorkuljis/green-tiles/assets/36756393/c1973071-c5cd-4630-91d3-c875fcafc2f5

## Notes

Here was the plan 🌏:
- Provide a `github username` as input and recieve `current years contribution graph` as output.

Problems 🫤:
- Github do not provide a REST api for this data over the web - you have to use a query client like GraphQL.

Solution 💭:
- Take a screenshot of github profile page and crop the contribution graph. Do this on a server and provide a web interface.

Obstacle-1 🗿:
- The contribution graph can be anywhere on the page, depending if the user has pinned repositories or a large README as a bio.
  - Fix: if a user wants to use the tool, they can simply remove the bio and select an offset option (1, 2 or 3) depending on how many pinned repositories they have.
  - Unexplored: Use computer-vision to detect a label, then draw a bounding box around that.

Obstacle-2 🗿:
- Then came deployment. I ssh'd into a digital ocean droplet, setup the project
and dependencies, only to find the chromium browser screeshots were totally
different to the one I was testing locally. At this point it came apparrent how
much a pain in the ass it would be to manage this - I decided to stop working on
the project.

Outcome 🐯:
- I can use this tool to capture screenshots for my own contribution graph, I run it on a defined schedule (eg: once per day). But not available to / for public.

> If trying to get your green squares over the internet was a poker game, I am folding; I'll let the GraphQL or the scripters win that pot.

- [https://github-contributions.vercel.app/](https://github-contributions.vercel.app/)



