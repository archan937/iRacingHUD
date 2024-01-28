window.JSX = `
// NAME:
// BOUNDS: 250,150 400x200

const Frame = styled.div\`
  position: absolute;
  width: 100%;
  height: 100%;
  opacity: 0.65;
  background: black;
  border-radius: 10px;
  z-index: 1;
\`;

const Flex = styled.div\`
  padding: 30px;
  display: flex;
  position: relative;
  z-index: 1;
\`;

const Label = styled.div\`
  color: white;
  font-family: Microsoft YaHei UI;
  font-size: 18px;
  font-weight: 400;
\`;

return (
  <>
    <Frame />
    <Flex>
      <div style={{ paddingRight: "12px" }}>
        <Label>Time:</Label>
        <Label>Laps:</Label>
        <Label>Drivers:</Label>
        <Label>SoF:</Label>
        <Label>Incidents:</Label>
      </div>
      <div>
        <Label>{time}</Label>
        <Label>
          {state.lap} / {state.totalLaps}
        </Label>
        <Label>{state.drivers}</Label>
        <Label>{state.sof}</Label>
        <Label>{state.incidents}</Label>
      </div>
    </Flex>
  </>
);
`;
