function uuid() {
  const chars = "0123456789abcdef";
  let hex = "";
  for (let i = 0; i < 8; ++i) {
    hex += chars.charAt(Math.floor(Math.random() * chars.length));
  }
  return hex;
}
