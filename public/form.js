export default class Form {
    /** @type {HTMLFontElement|null} */
    handler = null;
    /** @type {HTMLInputElement|null} */
    input = null;
    /** @type boolean */
    is_disabled = true;

    /**
     * @param {string} form_id - Form ID
     * @param {string} input_id - Input ID
     * */
    constructor(form_id, input_id) {
        this.handler = document.getElementById(form_id);
        this.input = document.getElementById(input_id);
        this.disable(this.is_disabled);
    }

    /** @param {boolean} disabled */
    disable(disabled) {
        this.is_disabled = disabled;
        for (const item of this.handler) {
            item.disabled = disabled;
        }
    }
}
