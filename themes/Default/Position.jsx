
// NAME: Position
// BOUNDS: 217,717 400x200

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

const Position = styled.div`
    color: white;
    font-family: Microsoft YaHei UI;
    font-size: 50px;
    font-weight: 400;
    line-height: 100%;
    text-align: center;
`;

const Competitor = styled(Position)`
    font-size: 20px;
`;

return (
    <>
    <Frame />
    <Flex>
        <Position>P{state.position}</Position>
        <div>
        <Competitor>
            ({state.ahead}) +{state.deltaAhead}
        </Competitor>
        <Competitor>
            ({state.behind}) -{state.deltaBehind}
        </Competitor>
        </div>
    </Flex>
    </>
);