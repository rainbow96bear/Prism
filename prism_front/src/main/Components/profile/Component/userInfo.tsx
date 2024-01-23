import { useNavigate, useParams } from "react-router-dom";
import { useEffect } from "react";
import styled from "styled-components";
import { useDispatch, useSelector } from "react-redux";

import { AppDispatch, RootState } from "../../../../app/store";
import { getPersonalDate } from "../../../../app/slices/profile/personal_data";
import ProfileImage from "../../../../CustomComponent/ProfileImg";

const UserInfo = () => {
  const dispatch = useDispatch<AppDispatch>();
  const personalDate = useSelector((state: RootState) => state.personal_data);
  const user = useSelector((state: RootState) => state.user);
  const navigator = useNavigate();

  const move = (path: string) => {
    navigator(path);
  };
  const { id } = useParams();
  useEffect(() => {
    dispatch(getPersonalDate(id));
    console.log(personalDate);
  }, []);

  return (
    <Container>
      <ProfileImage id={id != undefined ? id : "default"}></ProfileImage>
      <SubBox>
        <div className="nickname">{personalDate.nickname}</div>
        <div className="oneLineIntroduce">
          {personalDate.one_line_introduce == ""
            ? "한 줄 소개"
            : personalDate.one_line_introduce}
        </div>
        <HashtagBox>
          {personalDate?.hashtag?.map((value, index) => (
            <div className="hashtag" key={index}>
              #{value}
            </div>
          ))}
        </HashtagBox>
      </SubBox>
      <div>
        {user.user_id == id ? (
          <button
            onClick={() => {
              move("/profile/update/userinfo");
            }}>
            수정
          </button>
        ) : (
          <button>메세지</button>
        )}
      </div>
    </Container>
  );
};

export default UserInfo;

const Container = styled.div`
  display: flex;
  height: 200px;

  > img {
    border-radius: 100%;
    width: 200px;
  }
`;

const SubBox = styled.div`
  display: flex;
  flex-direction: column;
  justify-content: space-around;
  flex: 1;
  padding: 0px 20px;
  .nickname {
    font-weight: bold;
    font-size: 2rem;
  }
  .oneLineIntroduce {
    font-size: 1.2rem;
  }
  .hashtag {
    font-size: 1.2rem;
  }
`;

const HashtagBox = styled.div`
  display: flex;
  > div {
    padding: 0px 15px 0px 0px;
  }
`;
