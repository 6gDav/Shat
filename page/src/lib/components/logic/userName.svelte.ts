const isClient = typeof window !== 'undefined';

let initialName = "";
let initialIp = "";
if (isClient) {
  initialName = sessionStorage.getItem("userName") || "";
  initialIp = sessionStorage.getItem("userIp") || "";
}

let nameState = $state<string>(initialName);
let ipState = $state<string>(initialIp);

export const user = {
  //userName
  get userName() {
    return nameState;
  },
  set userName(value: string) {
    nameState = value;
    if (isClient) {
      sessionStorage.setItem("userName", value);
    }
  },
  //IP
  get initialIp() {
    return ipState;
  },
  set initialIp(value: string) {
    ipState = value;
    if (isClient) {
      sessionStorage.setItem("userIp", value);
    }
  }
};
