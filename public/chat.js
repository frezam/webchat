import Message from "./message.js";

export default class Chat {
    /** @type HTMLDivElement */
    handler = null;

    /** @param {string} id */
    constructor(id) {
        this.handler = document.getElementById(id);
    }

    /** @param {Message} message */
    append(message) {
        let msg = new Message(message);
        const $el = msg.mount();
        console.log($el);
        this.handler.append($el);
    }

    /** @param {string} id */
    remove(id) {
        for (const item of this.handler.querySelectorAll("")) {
        }
    }
}
