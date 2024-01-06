# Green Tiles

🔨 Automatic contribution graph screenshot tool. Enter any github username and return an image of your current contribution graph.

## Demo

https://github.com/connorkuljis/green-tiles/assets/36756393/c1973071-c5cd-4630-91d3-c875fcafc2f5

## Notes

Today is day 6/366 and represents milestone 🏆 for a full week of commits starting on the 31st! Thanks to Jason for creating this group and everyone who has been participating - it helps keep me accountable.

Here was the plan for the project I've been working on 🌏:
- Provide a `github username` as input and recieve `current years contribution graph` as output.

Problems 🫤:
- Github do not provide a REST api for this data over the web - you have to use a query client like GraphQL.

Solution 💭:
- Take a screenshot of github profile page and crop the contribution graph. Do this on a server and provide a web interface.

Obstacle-1 🗿:
- The contribution graph can be anywhere on the page, depending if the user has pinned repositories or a large README as a bio.
  - Fix: if a user wants to use the tool, they can simply remove the bio and select an offset option (1, 2 or 3) depending on how many pinned repositories they have or provide a custom offset.

Obstacle-2 🗿:
- Then came deployment. I ssh'd into a digital ocean droplet, setup the project
and dependencies, only to find the chromium browser screeshots were totally
different to the one I was testing locally. At this point it came apparrent how
much a pain in the ass it would be to manage this - I almost decided to stop working on
the project.
  - Fix: Containerised app with docker -> deploy using fly.io

Outcome 🐯:
- First MVP shipped of the year. Please try it out! A beta link is here:

https://green-tiles-young-field-7129.fly.dev/ 

> If trying to get your green squares over the internet was a poker game, I am folding; I'll let the GraphQL or the scripters win that pot.

- [https://github-contributions.vercel.app/](https://github-contributions.vercel.app/)



