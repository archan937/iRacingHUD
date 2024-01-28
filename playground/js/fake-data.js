window.updateFakeData = (setFakeData) => {
  const interval = setInterval(() => {
    setFakeData((draft) => {
      // Speed
    });
  }, Math.floor(Math.random() * (100 - 75 + 1) + 75));

  return () => {
    clearInterval(interval);
  };
};
