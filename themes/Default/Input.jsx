// NAME: Input
// BOUNDS: 437,1517 400x200

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
  padding: 20px;
  display: flex;
  position: relative;
  z-index: 1;
`;

const Gear = styled.div`
  color: white;
  font-family: Microsoft YaHei UI;
  font-size: 50px;
  font-weight: 400;
  line-height: 100%;
  text-align: center;
`;

const Unit = styled(Gear)`
  font-size: 20px;
`;

return (
  <>
    <Frame />
    <Flex>
      <Gear>{state.gear}</Gear>
      <div>
        <Unit>{state.speed}</Unit>
        <Unit>{state.rpm} rpm</Unit>
      </div>
    </Flex>
    <Flex>
      <div>
        <Unit>Throttle:</Unit>
        <Unit>Brake:</Unit>
        <Unit>Clutch:</Unit>
      </div>
      <div style={{ width: "400px" }}>
        <div
          style={{
            background: "green",
            width: `${state.throttle * 100}%`,
            height: "17px",
          }}
        >
          <Unit />
        </div>
        <div
          style={{
            background: "red",
            width: `${state.brake * 100}%`,
            height: "17px",
          }}
        >
          <Unit />
        </div>
        <div
          style={{
            background: "blue",
            width: `${(state.clutch == 1 ? 0 : state.clutch) * 100}%`,
            height: "17px",
          }}
        >
          <Unit />
        </div>
      </div>
    </Flex>
  </>
);
