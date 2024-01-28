
// NAME: Timing
// BOUNDS: 643,0 400x200

const Frame = styled.div`
    position: absolute;
    width: 100%;
    height: 100%;
    opacity: 0.65;
    background: black;
    border-radius: 10px;
    z-index: 1;
`;

const Flex = styled.div`
    padding: 30px;
    display: flex;
    position: relative;
    z-index: 1;
`;

const Time = styled.div`
    color: white;
    font-family: Microsoft YaHei UI;
    font-size: 50px;
    font-weight: 400;
    line-height: 100%;
    text-align: center;
`;

const Elapsed = styled(Time)`
    font-size: 20px;
`;

const Sector = styled(Elapsed)`
    maring-right: 2px;
    padding: 5px;
    color: black;
    flex: 1;
    font-size: 18px;
    font-weight: 800;
    background: grey;
`;

return (
    <>
    <Frame />
    <Flex>
        <Time>
        {state.sector === 1 || Object.keys(sectorTimes).length
            ? time
            : "00:00:00.000"}
        </Time>
        <div>
        <Elapsed>
            Last:{" "}
            {state.last !== -1
            ? moment.utc(state.last * 1000).format("HH:mm:ss.SSS")
            : ""}
        </Elapsed>
        <Elapsed>
            Best: {moment.utc(state.best * 1000).format("HH:mm:ss.SSS")}
        </Elapsed>
        </div>
    </Flex>
    <Flex style={{ marginTop: "-40px" }}>
        {Object.keys(sectorColors)
        .sort()
        .map((sector) => (
            <Sector
            key={`sector${sector}`}
            style={{
                color:
                { purple: "white" }[sectorColors[sector].color] ||
                "black",
                background: sectorColors[sector].color || "grey",
                opacity: sectorColors[sector].opacity || 0.75,
            }}
            >
            {sectorColors[sector].delta}
            </Sector>
        ))}
    </Flex>
    </>
);