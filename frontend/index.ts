import "htmx.org"
import htmx from "htmx.org"


interface ConfigRequestEvent extends Event {
	detail: {
		elt: Element,
		headers: Record<string, string>,
	}
}

interface SessionIdEvent extends Event {
	detail: {
		id: string
	}
} 

htmx.on('sessionId', function(evt: Event) {
   const sessionIdEvent = evt as SessionIdEvent
   const sessionId = sessionIdEvent.detail.id
	
   htmx.on('htmx:configRequest', function(evt: Event) {
	   const event = evt as ConfigRequestEvent
	   event.detail.headers["X-Session-Id"] =  sessionId
   })
})


