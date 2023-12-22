const app = (function () {
    const second = 1000
    const minute = second * 60

    let modalForm = null

    function getModal() {
        if (modalForm == null) {
            modalForm = new bootstrap.Modal(document.getElementById("modal-form"), {})
        }
        return modalForm
    }

    function openModal(scope, suggestedCommand, isNew) {
        const modalScope = document.getElementById("modal-scope")
        modalScope.value = scope
        modalScope.readOnly = !isNew
        onScopeChange(modalScope)

        const modalCommand = document.getElementById("modal-command")
        modalCommand.value = suggestedCommand

        const modalMinutes = document.getElementById("modal-minutes")
        modalMinutes.value = "60"

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

    function validateForm() {
        const modalScope = document.getElementById("modal-scope")
        if (!onScopeChange(modalScope)) {
            return false
        }
        const modalMinute = document.getElementById("modal-minutes")
        onMinuteChange(modalMinute)
        return true
    }

    return {
        nextInterval,
        openModal,
        cancel,
        validateForm,
    }
})()

function onMinuteChange(input) {
    if (!input.value) {
        input.value = 10
    }
    if (input.value < parseInt(10)) {
        input.value = 10
    }
}

function onScopeChange(input) {
    if (!input.value || input.value === "") {
        input.classList.add("is-invalid")
        return false
    } else {
        input.classList.remove("is-invalid")
        return true
    }
}