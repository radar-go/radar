# Radar
Radar is an implementation of the technology radar based on the one presented by [ThoughtWorks](https://www.thoughtworks.com/radar), which it's divided in four sections, *techniques*, *tools*, *platforms* and *language and frameworks*. For this project you can see three flavors of the radar, one for the experience of the people, another one with resources to learn about the technologies (books, videos, courses, ...) and a third one with projects done in the company. The purpose of this is not only to make visible the technologies used in your company, but to know the expertise of the people you're working with and to have resources to improve your team skills.

For now we're designing and implementing both a web interface and an API to access from other sources (like terminal commands, mobile applications, ...). Although initially it's a monolithic application, we're following the clean arquitecture and SOLID methodologies, so, if/when the program grow, we can think on split it in different services.

# How to use it

Both to build and test the project we make use of docker, using the `golang:1.9-alpine` image, that will be downloaded when you build the project or run the tests for the first time, so you need permissions to run docker containers on your development machine to work with the project. To build the **radar** command run:

```
make
```

and it will generate the binary in `bin/amd64/radar`. To run the tests you should run:

```
make tests
```

and it will show the tests results on the terminal. You can cross-compile **radar** by changing the variable ARCH in the Makefile by one of the next:

* amd64
* arm
* arm64
* ppc64le

# License
radar is licensed under the [GNU GPLv3](https://www.gnu.org/licenses/gpl.html). You should have received a copy of the GNU General Public License along with radar. If not, see http://www.gnu.org/licenses/.

<p align="center">
<img src="https://www.gnu.org/graphics/gplv3-127x51.png" alt="GNU GPLv3">
</p>
