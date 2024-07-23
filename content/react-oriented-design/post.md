---
title: React Oriented Design
publishDate: "2020-09-14"
path: react-oriented-design
---

Work on an older project has me building a lot of class components in React. Despite my affinity for object-oriented programming I've found the experience unsatisfying. Over the last month I've been trying to piece together why I feel this way.

Once upon a time class components were the only way to store state in React. As such they had very real use cases that could not be avoided. That changed with the release of hooks. Functional components now have just as much authority over state and lifecycle as their class counterparts. Still, functional and class components have the same capabilities so should developers have a preference for one over the other beyond mere asthetics? I think so, yes.

The case against class components might be thin, but it seems fundamentally sound to me. It's rooted in the idea of prefering [composition over inheritance][coi] in object-oriented design. When defining methods in class components those methods aren't part of the component's public interface. Other components do not send messages to this component corresponding to these methods. At best these methods should be considered in a protected scope (ES6 classes have no concept of member privacy) and used by components that extend it. Therefore it seems to me that methods in a class component barely function as methods at all. When they do it is in pursuit of a design concept best minimized in most codebases.

[coi]: https://en.wikipedia.org/wiki/Composition_over_inheritance
