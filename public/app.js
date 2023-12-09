import Form from "./form.js";
import Ws from "./ws.js";
import Chat from "./chat.js";

// TYPE DEFINITIONS
/**
 * @typedef {Object} User
 * @property {number} id
 * @property {string} name
 */

/**
 * @typedef {Object} Message
 * @property {number} id
 * @property {string} type
 * @property {User} user
 * @property {string} msg
 */

/**
 * @typedef {Object} WResponse
 * @property {boolean} error
 * @property {number} code
 * @property {string} message
 * @property {Message} data
 */

// TYPE DEFINITIONS

window.onload = () => {
    const form = new Form("form", "input")
        , ws = new Ws()
        , chat = new Chat("chat_controller");

    /* Websocket */
    ws.conn.onopen = (e) => {
        // form.disable(false);
        ws.conn.send("ping");
    }
    ws.conn.onmessage = (/** MessageEvent */e) => {
        /** @type WResponse */
        const response = JSON.parse(e?.data);
        if (response?.error) {
            console.error(response?.message);
            ws.conn.close();
            return
        }

        form.disable(false);
        form.input?.focus();
        const message = response?.data;
        chat.append(message);
    };

    /* FORM */
    form.handler.onsubmit = (e) => {
        e.preventDefault();
        const msg = form.input.value.trim(" ");
        if (msg === "") return false;

        ws.conn.send(msg);
        form.input.value = "";
        return false;
    };
};
