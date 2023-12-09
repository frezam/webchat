export default class Ws {
    /** @type WebSocket */
    conn = null;
    /** @type URLSearchParams */
    query = null;

    constructor() {
        this.query = new URLSearchParams(window.location.search);
        this.conn = new WebSocket(this.getUrl(
            document.location.host,
            this.query.get("id"),
            this.query.get("username"))
        );
    }

    /**
     * @param {string} host
     * @param {string} room_id
     * @param {string} username
     * @return {string} */
    getUrl(host, room_id, username) {
        return `ws://${host}/api/room?id=${room_id}&username=${username}`;
    }
}
