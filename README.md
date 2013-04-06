WALL-E
======

HOP THE FENCE!!!


This project implements an automatic solution to battle the Great FireWall (GFW). Using Go.

Current feature:
   * Automatically & periodically update hosts file, for those suffering from DNS poisoning.
   * see Roadmap below


## Compile and Usage ##

### Compile ###

#### smarthosts ####

go build smarthosts.go

###Usages ###
   * Under Windows: right-click the exe and select "Run as Administrator"
   * Under Mac/Linux: in terminal run "sudo ./smarthosts"


## Notes ##

Since Github.com is shutting down download section, the previous direct download of binaries will become unavailable.
From this point forward the download links would be pointed to file hosting services such as dropbox.com, and these will be
updated in wiki page [binary](../../wiki/Binary "Binary")

由于Github将不再提供文件下载（主要用来放可执行文件），往后的编译好的文件将在本项目的wiki的[binary](../../wiki/Binary "Binary")下面做更新。


## Roadmap for the project: (stand March 2013) ##

Rightnow the best bet for circumventing censorship is to use VPN service or SSH Tunneling, as they are most reliable. 
For people who are looking for reliable services, I would suggest using VPN services such as http://www.iqlinkus.com/products.action

This project, like many others, while trying its best to provide ease-of-use to every non-techie users, has the aim to showcase
the state of art effectiveness by researching how GFW works, and by direct countering its weakness. And such research requires:
  * sophisticated array of measurement to:
   * analyse the GFW working pattern
   * monitor and measure the countering effectiveness
   * provide proof to various concepts and assumptions

and

  * on-going commitment to:
   * study the mechanisms newly added to the GFW system
   * expose update-to-date insight of GFW development roadmaps
   * achieve certain degree of pressure in order to force GFW-development to raise its funding
   * motivate more and more people to join the good cause


Thus, my project will adapt the similar philosophy of project [West-Chamber](https://github.com/liruqi/west-chamber-season-3 "West-Chamber"). And here is the roadmap for the project:

  - [x] **provide some well-known, functional solution such as [smarthosts](https://code.google.com/p/smarthosts "SmartHosts").**
  - [ ] **provide ease-to-use setup for solutions like [goagent](https://code.google.com/p/goagent/ "GoAgent").**
  - [ ] **apply right tools(solution) for the right job(specific art of blockade) automatically**
   - [ ] **by probing/attacking GFW and figure out what kind of blockade is in place**
   - [ ] **and if available, choose the solution accordingly**


Also see wiki page [Roadmap](../../wiki/Roadmap "Roadmap")