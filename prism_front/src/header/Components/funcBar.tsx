import styled from "styled-components";
import { useEffect, useState } from "react";

import AfterLogin from "./beforeLogin";
import BeforeLogin from "./afterLogin";

const FuncBar: React.FC = () => {
  const [hasKakaoUserCookie, setHasKakaoUserCookie] = useState(false);

  useEffect(() => {
    const kakaoUserValue = document.cookie;
    const hasCookie = kakaoUserValue.includes("kakaoUser");

    setHasKakaoUserCookie(hasCookie);
  }, []);

  return <Box>{hasKakaoUserCookie ? <BeforeLogin /> : <AfterLogin />}</Box>;
};

export default FuncBar;

const Box = styled.div`
  display: flex;
  justify-content: right;
  height: 100%;
  align-items: center;
`;
