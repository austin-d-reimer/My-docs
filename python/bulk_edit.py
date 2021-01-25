def readFile(filename):
    lines = []
    with open(f"{filename}.md", "r", encoding="utf8") as f:
        for line in f:
            lines.append(line)
    return lines


def writeFile(filename, lines):
    with open(filename, "w", encoding="utf8") as f:
        f.writelines(lines)


def addBullets(filename):
    tweets = readFile(filename)
    editedTweets = []

    for tweet in tweets:
        editedTweets.append(f"* {tweet}")
    writeFile(f"{filename}.md", editedTweets)


def chronologicalOrder(filename):
    tweets = readFile(filename)
    editedTweets = tweets[-1:0:-1]
    writeFile(f"{filename}-chronoOrder.md", editedTweets)


def loadNew(filename):
    addBullets(filename)
    chronologicalOrder(filename)


loadNew("")
