$wordList = @('apple', 'banana', 'cherry', 'date', 'elderberry', 'fig', 'grape', 'honeydew', 'kiwi', 'lemon', 'mango', 'nectarine', 'orange', 'papaya', 'quince', 'raspberry', 'strawberry', 'tangerine', 'ugli', 'watermelon')

1..20 | ForEach-Object {
    $index = $_
    $randomWord = $wordList | Get-Random

    Write-Host "Round #$index"
    & .\leeter.exe -Input "$randomWord"
    Write-Host "------------------------------"
}
