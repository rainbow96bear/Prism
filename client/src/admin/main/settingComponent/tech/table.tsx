import React, { useState } from "react";
import styled from "styled-components";
import axios from "./../../../../configs/AxiosConfig";

interface TableProps {
  column: string[];
  info: any[][];
  getTechList: () => void;
}

const Table: React.FC<TableProps> = ({ column, info, getTechList }) => {
  const [editingRowIndex, setEditingRowIndex] = useState<number | null>(null);
  const [editedValues, setEditedValues] = useState<{ [key: string]: string }>(
    {}
  );

  const toggleEdit = (rowIndex: number) => {
    if (editingRowIndex === null) {
      setEditingRowIndex(rowIndex);
      // 기존 값으로 초기화
      setEditedValues(
        Object.fromEntries(
          column.map((key, index) => [key, info[rowIndex][index]])
        )
      );
    } else if (editingRowIndex === rowIndex) {
      setEditingRowIndex(null);
    }
  };
  const handleSave = async (rowIndex: number) => {
    try {
      // Axios로 서버에 저장 요청 보내기
      const result = (
        await axios.put(
          `/admins/techs/${rowIndex + 1}`,
          {
            tech_code: rowIndex + 1,
            tech_name: editedValues?.tech_name,
            count: 0,
          },
          {
            withCredentials: true,
          }
        )
      ).data;
      alert(
        "코드 번호 : " +
          result?.tech_code +
          ", 기술명 : " +
          result?.tech_name +
          "저장 완료"
      );
      // 서버에서 응답이 성공인 경우 상태 업데이트
      getTechList();
      const updatedInfo = [...info];
      updatedInfo[rowIndex] = Object.values(editedValues);
      setEditingRowIndex(null);
    } catch (error) {
      alert("중복된 코드를 입력하셨습니다.");
      console.error("Save failed:", error); // 에러 응답을 출력
      // 서버에서 응답이 실패한 경우에 대한 처리 추가
    }
  };

  const handleCancel = (rowIndex: number) => {
    // 취소 버튼 클릭 시 처리
    // 이 부분에서 편집 이전의 상태로 복구하는 로직을 추가하면 됩니다.
    setEditingRowIndex(null);
  };

  const handleChange = (
    e: React.ChangeEvent<HTMLInputElement>,
    key: string
  ) => {
    // input 값이 변경될 때마다 상태 업데이트
    setEditedValues((prevValues) => ({
      ...prevValues,
      [key]: e.target.value,
    }));
  };
  return (
    <TableContainer>
      <TableHeader>
        <TableRow>
          {column.map((columnName, index) => (
            <th key={index}>{columnName}</th>
          ))}
          <th></th>
        </TableRow>
      </TableHeader>
      <TableBody>
        {info.map((row, rowIndex) => (
          <TableRow key={rowIndex}>
            {Object.keys(row).map((key: any, cellIndex) => (
              <td key={cellIndex}>
                {editingRowIndex === rowIndex ? (
                  key === "tech_name" ? (
                    <input
                      type="text"
                      defaultValue={String(row[key])}
                      onChange={(e) => handleChange(e, key)}
                    />
                  ) : (
                    // Count는 입력란 대신 텍스트로 출력
                    row[key]
                  )
                ) : (
                  row[key]
                )}
              </td>
            ))}

            <td>
              {editingRowIndex === rowIndex ? (
                <>
                  <button onClick={() => handleSave(rowIndex)}>저장</button>
                  <button onClick={() => handleCancel(rowIndex)}>취소</button>
                </>
              ) : (
                <button onClick={() => toggleEdit(rowIndex)}>수정</button>
              )}
            </td>
          </TableRow>
        ))}
      </TableBody>
    </TableContainer>
  );
};

const TableContainer = styled.table`
  border-collapse: collapse;
  width: 100%;

  tr {
    > :first-child {
      border-left: none;
    }
    > :last-child {
      border-right: none;
    }
  }
`;

const TableHeader = styled.thead`
  th {
    border: 2px solid black;
    padding: 8px;
  }
`;

const TableBody = styled.tbody`
  td {
    border: 2px solid #ddd;
    padding: 10px;

    button {
      padding: 10px;
      border: 1px solid black;
      border-radius: 10px;
      cursor: pointer;
    }
  }
`;

const TableRow = styled.tr`
  height: 40px;

  td {
    border: 2px solid #ddd;
    padding: 10px;
    text-align: center; /* 텍스트를 가운데 정렬합니다. */
  }

  :last-child {
    td {
      width: 1%; /* fit-content를 대체하는 값으로 가변 너비를 가지게 합니다. */
    }
  }
`;

export default Table;
