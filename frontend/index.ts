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

interface DragEndEvent extends Event {
	x: number, 
	y: number,
}

class Square {
	elt: HTMLElement;
	boundingRect: DOMRect;
	constructor(elt: HTMLElement) {
		this.elt = elt
		this.boundingRect = elt.getBoundingClientRect()
	}
	
	replaceWith(elt: HTMLElement) {
		this.elt.id = elt.id
		this.elt.innerHTML = elt.innerHTML
		elt.id = ""
		elt.innerHTML = ""
	} 

}


class Board {
	container: HTMLElement
	squares: Square[]
	// squareMap: 

	constructor() {
		this.container = htmx.find("#board") as HTMLElement
		this.squares = (Array.from(htmx.findAll(this.container, ".board-square")) as HTMLElement[])
				     .map((elt: HTMLElement) => new Square(elt))
	}

	findSquare(x: number, y: number) {
		return this.squares.find(sq => {
			return sq.boundingRect.left < x  &&
			sq.boundingRect.right > x && 
			sq.boundingRect.top < y &&
			sq.boundingRect.bottom > y
		})
	}
}



htmx.onLoad(() => {
	
		

	htmx.findAll(".board-square").forEach(elt => {

		const board = new Board()
		elt.addEventListener('dragstart', () => {
			console.log("dragstart!")
		})
		elt.addEventListener('dragend', (e) => {
			const event = e as DragEndEvent
			const boundingRect = elt.getBoundingClientRect()
			console.log(boundingRect)
			console.log("dragend!")
			console.log(event)
			const newSquare = board.findSquare(event.x, event.y)
			newSquare?.replaceWith(elt as HTMLElement)
			console.log(newSquare)
		})
	})
})

// htmx.on('sessionId', function(evt: Event) {
//    const sessionIdEvent = evt as SessionIdEvent
//    const sessionId = sessionIdEvent.detail.id
//
//    htmx.on('htmx:configRequest', function(evt: Event) {
// 	   const event = evt as ConfigRequestEvent
// 	   event.detail.headers["X-Session-Id"] =  sessionId
//    })
// })


