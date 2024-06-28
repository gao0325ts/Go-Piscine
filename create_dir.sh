#!/bin/bash

if [ -z "$GO" ]; then
    echo "エラー: 環境変数GOが設定されていません。"
    exit 1
fi

if [ -z "$EX" ]; then
    echo "エラー: 環境変数EXが設定されていません。"
    exit 1
fi

project_num=${GO}
ex_count=${EX}

dir=$(printf "go%02d" $project_num)

if [ -d "$dir" ]; then
    read -p "ディレクトリ $dir は既に存在します。上書きしますか？ (y/n): " choice
    if [ "$choice" != "y" ]; then
        echo "処理を中止しました。"
        exit 1
    fi
    rm -rf "$dir"
fi

mkdir "$dir"

for i in $(seq 0 $ex_count)
do
    ex_dir=$(printf "ex%02d" $i)
    mkdir -p "$dir/$ex_dir/vendor/piscine"

    if [ -d "ft" ]; then
        cp -r ft "$dir/$ex_dir/vendor"
    else
        echo "警告: 'ft' ディレクトリが存在しません。コピーはスキップされました。" >&2
    fi
done

echo "ディレクトリ作成完了"