 
<h1><img src="https://rawgithub.com/runabove/metronome/master/icon.svg" width="48" height="48">&nbsp;Metronome - Distributed, fault tolerant scheduler</h1>

Metronome is a distributed and fault-tolerant event scheduler. It can be used to trigger remote systems throught events (HTTP, AMQP, KAFKA).

Metronome is written in Go and leverage the power of kafka streams to provide fault tolerance, reliability and scalability.

Metronome take a new approach to the job scheduling problem, as it only handle job scheduling not effective execution. Effective job execution is perform by triggered external system.

Metronome has a number of advantages over regular cron:
- Jobs can be written in any language, using any technology as it only trigger event.
- Jobs are schedule using [ISO8601][ISO8601] repeating interval notation, which enables more flexibility.
- It is able to handle high volumes of scheduled jobs in a completely fault way.
- Easy admin, thanks to a great [UI][UI].

## Status

Currently Metronome is under heavy development.

## Contributing

Instructions on how to contribute to Metronome are available on the [Contributing][Contributing] page.

## Get in touch

- Twitter: [@notd33d33](https://twitter.com/notd33d33)

[UI]: https://github.com/runabove/metronome-ui
[Contributing]: https://github.com/runabove/metronome/blob/master/README.md
[ISO8601]: http://en.wikipedia.org/wiki/ISO_8601 "ISO8601 Standard"
