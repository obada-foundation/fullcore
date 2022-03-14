export function copyToClipboard(str) {
    const el = document.createElement('textarea')
    el.value = str
    document.body.appendChild(el)
    el.select()
    el.setSelectionRange(0, 999999)
    document.execCommand('copy')
    document.body.removeChild(el)
}

export function shortenHash(str) {
    return str.substr(0, 10) + '...' + str.slice(-5)
}
