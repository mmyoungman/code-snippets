def makeBowlList(frames):
    bowlList = []
    frameIndex = {}
    for i, frame in enumerate(frames):
        frameIndex[i] = len(bowlList)
        if frame == "X":
            bowlList.append("X")
        elif len(frame) == 2:
            bowlList.append(frame[0])
            bowlList.append(frame[1])
        elif len(frame) == 3:
            bowlList.append(frame[0])
            bowlList.append(frame[1])
            bowlList.append(frame[2])

    for i, bowl in enumerate(bowlList):
        if bowl == "X":
            bowlList[i] = 10
        elif bowl == "/":
            bowlList[i] = 10 - bowlList[i-1]
        else:
            bowlList[i] = int(bowlList[i])

    return bowlList, frameIndex

def bowling_score(frames):
    bowlList, index = makeBowlList(frames.split())
    score = 0
    for i in range(10):
        if bowlList[ index[i] ] == 10 or bowlList[index[i]] + bowlList[index[i]+1] == 10:
            score += bowlList[index[i]] + bowlList[index[i]+1] + bowlList[index[i]+2]
        else:
            score += bowlList[index[i]] + bowlList[index[i]+1]
    return score