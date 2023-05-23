# Linux Container(LXC)

리눅스 상에서 사용하는 컨테이너 환경.

LXC는 리눅스 커널의 컨테이너 기능을 이용하기 위한 툴이나 API를 제공함.

컨테이너는 namespace와 cgroups라는 리소스 관리 장치를 사용하여 분리된 환경을 만듬.

데이터 영역에 대해서는 특정 디렉토리를 루트 디렉토리로 변경하는 chroot를 사용하여 분리 환경을 만듬.

Docker는 이전 버전에서는 내부에서 LXC를 사용했었지만 현재 버전에서는 사용하고 있지 않음.