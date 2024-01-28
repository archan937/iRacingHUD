function useImmerState() {
  const [getter, setter] = React.useState({});
  return [getter, (fn) => setter(immer.produce(fn))];
}

function syncGameState(setState, fn) {
  document.addEventListener("astilectron-ready", () => {
    astilectron.onMessage((json) => {
      const data = JSON.parse(json);
      setState(() => {
        return data;
      });
      fn && fn(data);
    });
  });
}
