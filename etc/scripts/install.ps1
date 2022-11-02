if (Get-Command "scoop" 2>$null) {
    scoop bucket add guardrails https://github.com/guardrailsio/scoop-bucket-guardrails
    scoop install guardrails
    scoop update guardrails
    Write-Output "GuardRails installed with Scoop! Run 'guardrails --help' to see available commands."
    return
} else {
    Write-Host "Scoop is not installed! (https://scoop.sh)"
}