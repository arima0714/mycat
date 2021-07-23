build:
	echo "docker build を実行する"
	docker build -t gopractice:latest ./

run:
	echo "bocker run を実行する"
	docker run -it --rm gopractice

buildWithNotOptimization:
	echo "docker build を実行する（非最適化）"
	docker build -f Dockerfile_not_opt -t gopractice_not_optimization:latest ./

runWithNotOptimization:
	echo "bocker run を実行する（非最適化）"
	docker run -it --rm gopractice_not_optimization
