const app = (function(){
    const second = 1000
    const minute = second * 60

    let modalForm = null

    function getModal() {
        if (modalForm == null) {
            modalForm = new bootstrap.Modal(document.getElementById("modal-form"), {})
        }
        return modalForm
    }

    function openModal(scope, suggestedCommand) {
        const modalScope = document.getElementById("modal-scope")
        modalScope.value = scope

        const modalCommand = document.getElementById("modal-command")
        modalCommand.value = suggestedCommand

        getModal().show()
    }

    function cancel() {
        getModal().hide()
    }

    function nextInterval(loadedAt, finishAt, elementName) {
        calculateTimeRemaining(loadedAt, finishAt, elementName)

        if (loadedAt > finishAt) {
            return
        }

        setTimeout(() => {
            nextInterval(loadedAt + second, finishAt, elementName)
        }, second)
    }

    function calculateTimeRemaining(startTime, endTime, elementName) {
        let timeRemaining = endTime - startTime
        if (timeRemaining < 0) {
            timeRemaining = 0
        }
        const remainingMinutes = `${Math.floor(timeRemaining / minute)}`.padStart(2, "0")
        const remainingSeconds = `${Math.floor((timeRemaining % minute) / second)}`.padStart(2, "0")

        document.getElementById(elementName).innerText = `Finish in ${remainingMinutes}:${remainingSeconds}`
    }
    return {
        nextInterval,
        openModal,
        cancel
    }
})()
