export class ProjectState {
    button_state: "enabled" | "a" | "b" = $state("enabled")

    select(option: "a" | "b") {
        console.log(option)
        if (this.button_state == "enabled") {
            this.button_state = option
        }
    }

    enable() {
        this.button_state = "enabled"
    }
}
