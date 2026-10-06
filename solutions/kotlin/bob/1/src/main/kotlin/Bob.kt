object Bob {
    fun hey(input: String): String {
        val trimmed = input.trim()
        val isSilence = trimmed.isEmpty()
        val isQuestion = trimmed.endsWith("?")
        val isYell = trimmed.any { it.isLetter() } && trimmed.none { it.isLowerCase() }

        return when {
            isSilence -> "Fine. Be that way!"
            isYell && isQuestion -> "Calm down, I know what I'm doing!"
            isYell -> "Whoa, chill out!"
            isQuestion -> "Sure."
            else -> "Whatever."
        }
    }
}