/**
 * @param {string} tagName
 * @param {{classList, id}} args
 * @return HTMLDivElement
 */
const cE = (tagName, args) => {
    const $el = document.createElement(tagName);
    for (const prop in args) {
        $el[prop] = args[prop]
    }
    return $el;
}

export default class Message {
    /** @type {Message|null} */
    message = null;

    /** @param {Message} message */
    constructor(message) {
        this.message = message;
    }

    /** @return HTMLDivElement */
    mount() {
        const $el = cE("div", {classList: "message-row"});
        const $message = cE("div", {classList: "text"});

        if (this.message.type === "system") {
            $el.classList.add("message-system");
            $message.textContent = this.message?.msg;

        } else if (this.message?.type === "user") {
            $el.classList.add("message-user");
            $message.textContent = `${this.message.user.name}: ${this.message.msg}`;

        } else {
            console.error(this.message);
        }

        $el.append($message);

        return $el;
    }
}
