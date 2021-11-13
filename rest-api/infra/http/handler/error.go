package handler

// BadRequestError 는 http 핸들러가 인자나 바디 등을 parsing하는 중에 에러가 발생하거나 하는 등의 에러
// 리턴할 때 해당 타입으로 감싸서 리턴하고 errors.As로 판단하는 방식
type BadRequestError struct {
	error
}

// variable인 에러를 감싸서 리턴하고 errors.Is로 판단하는 방식
//var (
//	BadRequestError = goerrors.New("잘못된 요청입니다.")
//)

// 뭐가 나을까요 ㅋㅋㅋ ㅜㅜㅜ
