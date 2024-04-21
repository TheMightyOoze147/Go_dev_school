# Построение AST
<details>
    ```console
    before walk main
    .   DCL # main.go:4:2
    .   .   NAME-main.a esc(no) Class:PAUTO Offset:0 OnStack Used int tc(1) # main.go:4:2
    .   AS Def tc(1) # main.go:4:4
    .   .   NAME-main.a esc(no) Class:PAUTO Offset:0 OnStack Used int tc(1) # main.go:4:2
    .   .   LITERAL-1 int tc(1) # main.go:4:7
    .   DCL # main.go:5:2
    .   .   NAME-main.b esc(no) Class:PAUTO Offset:0 OnStack Used int tc(1) # main.go:5:2
    .   AS Def tc(1) # main.go:5:4
    .   .   NAME-main.b esc(no) Class:PAUTO Offset:0 OnStack Used int tc(1) # main.go:5:2
    .   .   LITERAL-2 int tc(1) # main.go:5:7
    .   DCL # main.go:7:6
    .   .   NAME-main.a esc(no) Class:PAUTO Offset:0 InlFormal OnStack Used int tc(1) # main.go:7:6 main.go:11:10
    .   DCL # main.go:7:6
    .   .   NAME-main.b esc(no) Class:PAUTO Offset:0 InlFormal OnStack Used int tc(1) # main.go:7:6 main.go:11:13
    .   AS2 Def tc(1) # main.go:7:6
    .   AS2-Lhs
    .   .   NAME-main.a esc(no) Class:PAUTO Offset:0 InlFormal OnStack Used int tc(1) # main.go:7:6 main.go:11:10
    .   .   NAME-main.b esc(no) Class:PAUTO Offset:0 InlFormal OnStack Used int tc(1) # main.go:7:6 main.go:11:13
    .   AS2-Rhs
    .   .   NAME-main.a esc(no) Class:PAUTO Offset:0 OnStack Used int tc(1) # main.go:4:2
    .   .   NAME-main.b esc(no) Class:PAUTO Offset:0 OnStack Used int tc(1) # main.go:5:2
    .   INLMARK # +main.go:7:6
    .   PRINTN tc(1) # main.go:7:6 main.go:12:9
    .   PRINTN-Args
    .   .   ADD int tc(1) # main.go:7:6 main.go:12:12
    .   .   .   NAME-main.a esc(no) Class:PAUTO Offset:0 InlFormal OnStack Used int tc(1) # main.go:7:6 main.go:11:10
    .   .   .   NAME-main.b esc(no) Class:PAUTO Offset:0 InlFormal OnStack Used int tc(1) # main.go:7:6 main.go:11:13
    .   LABEL main..i0 # main.go:7:6
    after walk main
    .   DCL # main.go:4:2
    .   .   NAME-main.a esc(no) Class:PAUTO Offset:0 OnStack Used int tc(1) # main.go:4:2
    .   AS Def tc(1) # main.go:4:4
    .   .   NAME-main.a esc(no) Class:PAUTO Offset:0 OnStack Used int tc(1) # main.go:4:2
    .   .   LITERAL-1 int tc(1) # main.go:4:7
    .   DCL # main.go:5:2
    .   .   NAME-main.b esc(no) Class:PAUTO Offset:0 OnStack Used int tc(1) # main.go:5:2
    .   AS Def tc(1) # main.go:5:4
    .   .   NAME-main.b esc(no) Class:PAUTO Offset:0 OnStack Used int tc(1) # main.go:5:2
    .   .   LITERAL-2 int tc(1) # main.go:5:7
    .   DCL # main.go:7:6
    .   .   NAME-main.a esc(no) Class:PAUTO Offset:0 InlFormal OnStack Used int tc(1) # main.go:7:6 main.go:11:10
    .   DCL # main.go:7:6
    .   .   NAME-main.b esc(no) Class:PAUTO Offset:0 InlFormal OnStack Used int tc(1) # main.go:7:6 main.go:11:13
    .   BLOCK # main.go:7:6
    .   BLOCK-List
    .   .   AS tc(1) # main.go:7:6
    .   .   .   NAME-main.a esc(no) Class:PAUTO Offset:0 InlFormal OnStack Used int tc(1) # main.go:7:6 main.go:11:10
    .   .   .   NAME-main.a esc(no) Class:PAUTO Offset:0 OnStack Used int tc(1) # main.go:4:2
    .   .   AS tc(1) # main.go:7:6
    .   .   .   NAME-main.b esc(no) Class:PAUTO Offset:0 InlFormal OnStack Used int tc(1) # main.go:7:6 main.go:11:13
    .   .   .   NAME-main.b esc(no) Class:PAUTO Offset:0 OnStack Used int tc(1) # main.go:5:2
    .   INLMARK # +main.go:7:6
    .   BLOCK-init
    .   .   AS tc(1) # main.go:7:6 main.go:12:9
    .   .   .   NAME-main..autotmp_4 esc(N) Class:PAUTO Offset:0 AutoTemp OnStack Used int tc(1) # main.go:7:6 main.go:12:9
    .   .   .   ADD int tc(1) # main.go:7:6 main.go:12:12
    .   .   .   .   NAME-main.a esc(no) Class:PAUTO Offset:0 InlFormal OnStack Used int tc(1) # main.go:7:6 main.go:11:10
    .   .   .   .   NAME-main.b esc(no) Class:PAUTO Offset:0 InlFormal OnStack Used int tc(1) # main.go:7:6 main.go:11:13
    .   BLOCK tc(1) # main.go:7:6 main.go:12:9
    .   BLOCK-List
    .   .   CALLFUNC Walked tc(1) # main.go:7:6 main.go:12:9
    .   .   CALLFUNC-Fun
    .   .   .   NAME-runtime.printlock Class:PFUNC Offset:0 Used FUNC-func() tc(1)
    .   .   CALLFUNC Walked tc(1) # main.go:7:6 main.go:12:9
    .   .   CALLFUNC-Fun
    .   .   .   NAME-runtime.printint Class:PFUNC Offset:0 Used FUNC-func(int64) tc(1)
    .   .   CALLFUNC-Args
    .   .   .   CONVNOP int64 tc(1) # main.go:7:6 main.go:12:9
    .   .   .   .   NAME-main..autotmp_4 esc(N) Class:PAUTO Offset:0 AutoTemp OnStack Used int tc(1) # main.go:7:6 main.go:12:9
    .   .   CALLFUNC Walked tc(1) # main.go:7:6 main.go:12:9
    .   .   CALLFUNC-Fun
    .   .   .   NAME-runtime.printnl Class:PFUNC Offset:0 Used FUNC-func() tc(1)
    .   .   CALLFUNC Walked tc(1) # main.go:7:6 main.go:12:9
    .   .   CALLFUNC-Fun
    .   .   .   NAME-runtime.printunlock Class:PFUNC Offset:0 Used FUNC-func() tc(1)
    .   LABEL main..i0 # main.go:7:6

    before walk add
    .   PRINTN tc(1) # main.go:12:9
    .   PRINTN-Args
    .   .   ADD int tc(1) # main.go:12:12
    .   .   .   NAME-main.a esc(no) Class:PPARAM Offset:0 OnStack Used int tc(1) # main.go:11:10      
    .   .   .   NAME-main.b esc(no) Class:PPARAM Offset:0 OnStack Used int tc(1) # main.go:11:13      
    after walk add
    .   BLOCK-init
    .   .   AS tc(1) # main.go:12:9
    .   .   .   NAME-main..autotmp_2 esc(N) Class:PAUTO Offset:0 AutoTemp OnStack Used int tc(1) # main.go:12:9
    .   .   .   ADD int tc(1) # main.go:12:12
    .   .   .   .   NAME-main.a esc(no) Class:PPARAM Offset:0 OnStack Used int tc(1) # main.go:11:10  
    .   .   .   .   NAME-main.b esc(no) Class:PPARAM Offset:0 OnStack Used int tc(1) # main.go:11:13  
    .   BLOCK tc(1) # main.go:12:9
    .   BLOCK-List
    .   .   CALLFUNC Walked tc(1) # main.go:12:9
    .   .   CALLFUNC-Fun
    .   .   .   NAME-runtime.printlock Class:PFUNC Offset:0 Used FUNC-func() tc(1)
    .   .   CALLFUNC Walked tc(1) # main.go:12:9
    .   .   CALLFUNC-Fun
    .   .   .   NAME-runtime.printint Class:PFUNC Offset:0 Used FUNC-func(int64) tc(1)
    .   .   CALLFUNC-Args
    .   .   .   CONVNOP int64 tc(1) # main.go:12:9
    .   .   .   .   NAME-main..autotmp_2 esc(N) Class:PAUTO Offset:0 AutoTemp OnStack Used int tc(1) # main.go:12:9
    .   .   CALLFUNC Walked tc(1) # main.go:12:9
    .   .   CALLFUNC-Fun
    .   .   .   NAME-runtime.printnl Class:PFUNC Offset:0 Used FUNC-func() tc(1)
    .   .   CALLFUNC Walked tc(1) # main.go:12:9
    .   .   CALLFUNC-Fun
    .   .   .   NAME-runtime.printunlock Class:PFUNC Offset:0 Used FUNC-func() tc(1)
    ```
</details>

# Построение SSA
## На операционной системе Windows 10 не удалось с помощью команды GOSSAFUNC=main go tool compile ./main.go > ssa.html построить SSA, поэтому SSA был построен через сервис https://golang.design/gossa
### SSA для main()
<details>
    ```console
    AST
    buildssa-enter
    buildssa-body
    . DCL # main.go:4:2
    . . NAME-main.a esc(no) Class:PAUTO Offset:0 OnStack Used int tc(1) # main.go:4:2
    . AS Def tc(1) # main.go:4:4
    . . NAME-main.a esc(no) Class:PAUTO Offset:0 OnStack Used int tc(1) # main.go:4:2
    . . LITERAL-1 int tc(1) # main.go:4:7
    . DCL # main.go:5:2
    . . NAME-main.b esc(no) Class:PAUTO Offset:0 OnStack Used int tc(1) # main.go:5:2
    . AS Def tc(1) # main.go:5:4
    . . NAME-main.b esc(no) Class:PAUTO Offset:0 OnStack Used int tc(1) # main.go:5:2
    . . LITERAL-2 int tc(1) # main.go:5:7
    . IF tc(1) # main.go:6:2
    . IF-Cond
    . . LITERAL-true bool tc(1) # main.go:6:5
    . IF-Body
    . . DCL # main.go:7:6
    . . . NAME-main.a esc(no) Class:PAUTO Offset:0 InlFormal OnStack Used int tc(1) # main.go:7:6 main.go:11:10
    . . DCL # main.go:7:6
    . . . NAME-main.b esc(no) Class:PAUTO Offset:0 InlFormal OnStack Used int tc(1) # main.go:7:6 main.go:11:13
    . . BLOCK # main.go:7:6
    . . BLOCK-List
    . . . AS tc(1) # main.go:7:6
    . . . . NAME-main.a esc(no) Class:PAUTO Offset:0 InlFormal OnStack Used int tc(1) # main.go:7:6 main.go:11:10
    . . . . NAME-main.a esc(no) Class:PAUTO Offset:0 OnStack Used int tc(1) # main.go:4:2
    . . . AS tc(1) # main.go:7:6
    . . . . NAME-main.b esc(no) Class:PAUTO Offset:0 InlFormal OnStack Used int tc(1) # main.go:7:6 main.go:11:13
    . . . . NAME-main.b esc(no) Class:PAUTO Offset:0 OnStack Used int tc(1) # main.go:5:2
    . . INLMARK # +main.go:7:6
    . . BLOCK-init
    . . . AS tc(1) # main.go:7:6 main.go:12:9
    . . . . NAME-main..autotmp_4 esc(N) Class:PAUTO Offset:0 AutoTemp OnStack Used int tc(1) # main.go:7:6 main.go:12:9
    . . . . ADD int tc(1) # main.go:7:6 main.go:12:12
    . . . . . NAME-main.a esc(no) Class:PAUTO Offset:0 InlFormal OnStack Used int tc(1) # main.go:7:6 main.go:11:10
    . . . . . NAME-main.b esc(no) Class:PAUTO Offset:0 InlFormal OnStack Used int tc(1) # main.go:7:6 main.go:11:13
    . . BLOCK tc(1) # main.go:7:6 main.go:12:9
    . . BLOCK-List
    . . . CALLFUNC Walked tc(1) # main.go:7:6 main.go:12:9
    . . . . NAME-runtime.printlock Class:PFUNC Offset:0 Used FUNC-func() tc(1)
    . . . CALLFUNC Walked tc(1) # main.go:7:6 main.go:12:9
    . . . . NAME-runtime.printint Class:PFUNC Offset:0 Used FUNC-func(int64) tc(1)
    . . . CALLFUNC-Args
    . . . . CONVNOP int64 tc(1) # main.go:7:6 main.go:12:9
    . . . . . NAME-main..autotmp_4 esc(N) Class:PAUTO Offset:0 AutoTemp OnStack Used int tc(1) # main.go:7:6 main.go:12:9
    . . . CALLFUNC Walked tc(1) # main.go:7:6 main.go:12:9
    . . . . NAME-runtime.printnl Class:PFUNC Offset:0 Used FUNC-func() tc(1)
    . . . CALLFUNC Walked tc(1) # main.go:7:6 main.go:12:9
    . . . . NAME-runtime.printunlock Class:PFUNC Offset:0 Used FUNC-func() tc(1)
    . . LABEL main..i0 # main.go:7:6
    buildssa-exit
    
    before insert phis
    b1:-
    v1 (?) = InitMem <mem>
    v2 (?) = SP <uintptr>
    v3 (?) = SB <uintptr>
    v4 (?) = Const64 <int> [1] (a[int], a[int])
    v5 (?) = Const64 <int> [2] (b[int], b[int])
    v6 (+7) = InlMark <void> [0] v1
    v7 (12) = Add64 <int> v4 v5
    v8 (12) = StaticLECall <mem> {AuxCall{runtime.printlock}} v1
    v9 (12) = SelectN <mem> [0] v8
    v10 (12) = Copy <int64> v7
    v11 (12) = StaticLECall <mem> {AuxCall{runtime.printint}} [8] v10 v9
    v12 (12) = SelectN <mem> [0] v11
    v13 (12) = StaticLECall <mem> {AuxCall{runtime.printnl}} v12
    v14 (12) = SelectN <mem> [0] v13
    v15 (12) = StaticLECall <mem> {AuxCall{runtime.printunlock}} v14
    v16 (12) = SelectN <mem> [0] v15
    Plain → b2 (7)
    b2: ← b1-
    v18 (9) = FwdRef <mem> {{[] mem}}
    v17 (9) = MakeResult <mem> v18
    Ret v17 (9)
    name a[int]: v4
    name b[int]: v5
    name a[int]: v4
    name b[int]: v5

    start
    b1:-
    v1 (?) = InitMem <mem>
    v2 (?) = SP <uintptr>
    v3 (?) = SB <uintptr>
    v4 (?) = Const64 <int> [1] (a[int], a[int])
    v5 (?) = Const64 <int> [2] (b[int], b[int])
    v6 (+7) = InlMark <void> [0] v1
    v7 (12) = Add64 <int> v4 v5
    v8 (12) = StaticLECall <mem> {AuxCall{runtime.printlock}} v1
    v9 (12) = SelectN <mem> [0] v8
    v10 (12) = Copy <int64> v7
    v11 (12) = StaticLECall <mem> {AuxCall{runtime.printint}} [8] v10 v9
    v12 (12) = SelectN <mem> [0] v11
    v13 (12) = StaticLECall <mem> {AuxCall{runtime.printnl}} v12
    v14 (12) = SelectN <mem> [0] v13
    v15 (12) = StaticLECall <mem> {AuxCall{runtime.printunlock}} v14
    v16 (12) = SelectN <mem> [0] v15
    Plain → b2 (7)
    b2: ← b1-
    v18 (9) = Copy <mem> v16
    v17 (9) = MakeResult <mem> v18
    Ret v17 (9)
    name a[int]: v4
    name b[int]: v5
    name a[int]: v4
    name b[int]: v5

    number lines [9685 ns]
    b1:-
    v1 (?) = InitMem <mem>
    v2 (?) = SP <uintptr>
    v3 (?) = SB <uintptr>
    v4 (?) = Const64 <int> [1] (a[int], a[int])
    v5 (?) = Const64 <int> [2] (b[int], b[int])
    v6 (+7) = InlMark <void> [0] v1
    v7 (+12) = Add64 <int> v4 v5
    v8 (12) = StaticLECall <mem> {AuxCall{runtime.printlock}} v1
    v9 (12) = SelectN <mem> [0] v8
    v10 (12) = Copy <int64> v7
    v11 (12) = StaticLECall <mem> {AuxCall{runtime.printint}} [8] v10 v9
    v12 (12) = SelectN <mem> [0] v11
    v13 (12) = StaticLECall <mem> {AuxCall{runtime.printnl}} v12
    v14 (12) = SelectN <mem> [0] v13
    v15 (12) = StaticLECall <mem> {AuxCall{runtime.printunlock}} v14
    v16 (12) = SelectN <mem> [0] v15
    Plain → b2 (+7)
    b2: ← b1-
    v18 (9) = Copy <mem> v16
    v17 (+9) = MakeResult <mem> v18
    Ret v17 (9)
    name a[int]: v4
    name b[int]: v5
    name a[int]: v4
    name b[int]: v5

    early phielim [1298 ns]
    b1:-
    v1 (?) = InitMem <mem>
    v2 (?) = SP <uintptr>
    v3 (?) = SB <uintptr>
    v4 (?) = Const64 <int> [1] (a[int], a[int])
    v5 (?) = Const64 <int> [2] (b[int], b[int])
    v6 (+7) = InlMark <void> [0] v1
    v7 (+12) = Add64 <int> v4 v5
    v8 (12) = StaticLECall <mem> {AuxCall{runtime.printlock}} v1
    v9 (12) = SelectN <mem> [0] v8
    v10 (12) = Copy <int64> v7
    v11 (12) = StaticLECall <mem> {AuxCall{runtime.printint}} [8] v7 v9
    v12 (12) = SelectN <mem> [0] v11
    v13 (12) = StaticLECall <mem> {AuxCall{runtime.printnl}} v12
    v14 (12) = SelectN <mem> [0] v13
    v15 (12) = StaticLECall <mem> {AuxCall{runtime.printunlock}} v14
    v16 (12) = SelectN <mem> [0] v15
    Plain → b2 (+7)
    b2: ← b1-
    v17 (+9) = MakeResult <mem> v16
    v18 (9) = Copy <mem> v16
    Ret v17 (9)
    name a[int]: v4
    name b[int]: v5
    name a[int]: v4
    name b[int]: v5

    early copyelim [830 ns]
    early deadcode [7438 ns]
    b1:-
    v1 (?) = InitMem <mem>
    v4 (?) = Const64 <int> [1] (a[int], a[int])
    v5 (?) = Const64 <int> [2] (b[int], b[int])
    v6 (+7) = InlMark <void> [0] v1
    v7 (+12) = Add64 <int> v4 v5
    v8 (12) = StaticLECall <mem> {AuxCall{runtime.printlock}} v1
    v9 (12) = SelectN <mem> [0] v8
    v11 (12) = StaticLECall <mem> {AuxCall{runtime.printint}} [8] v7 v9
    v12 (12) = SelectN <mem> [0] v11
    v13 (12) = StaticLECall <mem> {AuxCall{runtime.printnl}} v12
    v14 (12) = SelectN <mem> [0] v13
    v15 (12) = StaticLECall <mem> {AuxCall{runtime.printunlock}} v14
    v16 (12) = SelectN <mem> [0] v15
    Plain → b2 (+7)
    b2: ← b1-
    v17 (+9) = MakeResult <mem> v16
    Ret v17 (9)
    name a[int]: v4
    name b[int]: v5
    name a[int]: v4
    name b[int]: v5

    short circuit [6901 ns]
    b1:
    BlockInvalid (+7)
    b2:-
    v1 (?) = InitMem <mem>
    v4 (?) = Const64 <int> [1] (a[int], a[int])
    v5 (?) = Const64 <int> [2] (b[int], b[int])
    v6 (+7) = InlMark <void> [0] v1
    v7 (+12) = Add64 <int> v4 v5
    v8 (12) = StaticLECall <mem> {AuxCall{runtime.printlock}} v1
    v9 (12) = SelectN <mem> [0] v8
    v11 (12) = StaticLECall <mem> {AuxCall{runtime.printint}} [8] v7 v9
    v12 (12) = SelectN <mem> [0] v11
    v13 (12) = StaticLECall <mem> {AuxCall{runtime.printnl}} v12
    v14 (12) = SelectN <mem> [0] v13
    v15 (12) = StaticLECall <mem> {AuxCall{runtime.printunlock}} v14
    v16 (12) = SelectN <mem> [0] v15
    v17 (+9) = MakeResult <mem> v16
    Ret v17 (9)
    name a[int]: v4
    name b[int]: v5
    name a[int]: v4
    name b[int]: v5

    decompose user [741 ns]
    pre-opt deadcode [3543 ns]
    b2:-
    v1 (?) = InitMem <mem>
    v4 (?) = Const64 <int> [1] (a[int], a[int])
    v5 (?) = Const64 <int> [2] (b[int], b[int])
    v6 (+7) = InlMark <void> [0] v1
    v7 (+12) = Add64 <int> v4 v5
    v8 (12) = StaticLECall <mem> {AuxCall{runtime.printlock}} v1
    v9 (12) = SelectN <mem> [0] v8
    v11 (12) = StaticLECall <mem> {AuxCall{runtime.printint}} [8] v7 v9
    v12 (12) = SelectN <mem> [0] v11
    v13 (12) = StaticLECall <mem> {AuxCall{runtime.printnl}} v12
    v14 (12) = SelectN <mem> [0] v13
    v15 (12) = StaticLECall <mem> {AuxCall{runtime.printunlock}} v14
    v16 (12) = SelectN <mem> [0] v15
    v17 (+9) = MakeResult <mem> v16
    Ret v17 (9)
    name a[int]: v4
    name b[int]: v5
    name a[int]: v4
    name b[int]: v5

    opt [33603 ns]
    b2:-
    v1 (?) = InitMem <mem>
    v6 (+7) = InlMark <void> [0] v1
    v7 (12) = Const64 <int> [3]
    v8 (+12) = StaticLECall <mem> {AuxCall{runtime.printlock}} v1
    v9 (12) = SelectN <mem> [0] v8
    v11 (12) = StaticLECall <mem> {AuxCall{runtime.printint}} [8] v7 v9
    v12 (12) = SelectN <mem> [0] v11
    v13 (12) = StaticLECall <mem> {AuxCall{runtime.printnl}} v12
    v14 (12) = SelectN <mem> [0] v13
    v15 (12) = StaticLECall <mem> {AuxCall{runtime.printunlock}} v14
    v16 (12) = SelectN <mem> [0] v15
    v17 (+9) = MakeResult <mem> v16
    Ret v17 (9)
    name a[int]: v4
    name b[int]: v5
    name a[int]: v4
    name b[int]: v5

    zero arg cse [2522 ns]
    opt deadcode [2796 ns]
    b2:-
    v1 (?) = InitMem <mem>
    v6 (+7) = InlMark <void> [0] v1
    v7 (12) = Const64 <int> [3]
    v8 (+12) = StaticLECall <mem> {AuxCall{runtime.printlock}} v1
    v9 (12) = SelectN <mem> [0] v8
    v11 (12) = StaticLECall <mem> {AuxCall{runtime.printint}} [8] v7 v9
    v12 (12) = SelectN <mem> [0] v11
    v13 (12) = StaticLECall <mem> {AuxCall{runtime.printnl}} v12
    v14 (12) = SelectN <mem> [0] v13
    v15 (12) = StaticLECall <mem> {AuxCall{runtime.printunlock}} v14
    v16 (12) = SelectN <mem> [0] v15
    v17 (+9) = MakeResult <mem> v16
    Ret v17 (9)

    generic cse [11955 ns]
    phiopt [676 ns]
    gcse deadcode [1559 ns]
    nilcheckelim [4248 ns]
    prove [13523 ns]
    b2:-
    v1 (?) = InitMem <mem>
    v6 (+7) = InlMark <void> [0] v1
    v7 (12) = Const64 <int> [3]
    v8 (+12) = StaticLECall <mem> {AuxCall{runtime.printlock}} v1
    v9 (12) = SelectN <mem> [0] v8
    v11 (12) = StaticLECall <mem> {AuxCall{runtime.printint}} [8] v7 v9
    v12 (12) = SelectN <mem> [0] v11
    v13 (12) = StaticLECall <mem> {AuxCall{runtime.printnl}} v12
    v14 (12) = SelectN <mem> [0] v13
    v15 (12) = StaticLECall <mem> {AuxCall{runtime.printunlock}} v14
    v16 (12) = SelectN <mem> [0] v15
    v17 (+9) = MakeResult <mem> v16
    v5 (?) = Const64 <int64> [0]
    Ret v17 (9)

    early fuse [421 ns]
    decompose builtin [6587 ns]
    expand calls [12092 ns]
    b2:-
    v1 (?) = InitMem <mem>
    v6 (+7) = InlMark <void> [0] v1
    v7 (12) = Const64 <int> [3]
    v8 (+12) = StaticCall <mem> {AuxCall{runtime.printlock}} v1
    v9 (12) = SelectN <mem> [0] v8
    v11 (12) = StaticCall <mem> {AuxCall{runtime.printint}} [8] v7 v9
    v12 (12) = SelectN <mem> [0] v11
    v13 (12) = StaticCall <mem> {AuxCall{runtime.printnl}} v12
    v14 (12) = SelectN <mem> [0] v13
    v15 (12) = StaticCall <mem> {AuxCall{runtime.printunlock}} v14
    v16 (12) = SelectN <mem> [0] v15
    v17 (+9) = MakeResult <mem> v16
    v5 (?) = Const64 <int64> [0]
    v4 (?) = SB <uintptr>
    v18 (?) = SP <uintptr>
    Ret v17 (9)

    softfloat [197 ns]
    late opt [2348 ns]
    b2:-
    v1 (?) = InitMem <mem>
    v6 (+7) = InlMark <void> [0] v1
    v7 (12) = Const64 <int> [3]
    v8 (+12) = StaticCall <mem> {AuxCall{runtime.printlock}} v1
    v9 (12) = SelectN <mem> [0] v8
    v11 (12) = StaticCall <mem> {AuxCall{runtime.printint}} [8] v7 v9
    v12 (12) = SelectN <mem> [0] v11
    v13 (12) = StaticCall <mem> {AuxCall{runtime.printnl}} v12
    v14 (12) = SelectN <mem> [0] v13
    v15 (12) = StaticCall <mem> {AuxCall{runtime.printunlock}} v14
    v16 (12) = SelectN <mem> [0] v15
    v17 (+9) = MakeResult <mem> v16
    Ret v17 (9)

    dead auto elim [1273 ns]
    generic deadcode [1773 ns]
    check bce [184 ns]
    branchelim [1309 ns]
    late fuse [2186 ns]
    b2:-
    v1 (?) = InitMem <mem>
    v6 (+7) = InlMark <void> [0] v1
    v7 (12) = Const64 <int> [3]
    v8 (+12) = StaticCall <mem> {AuxCall{runtime.printlock}} v1
    v9 (12) = SelectN <mem> [0] v8
    v11 (12) = StaticCall <mem> {AuxCall{runtime.printint}} [8] v7 v9
    v12 (12) = SelectN <mem> [0] v11
    v13 (12) = StaticCall <mem> {AuxCall{runtime.printnl}} v12
    v14 (12) = SelectN <mem> [0] v13
    v15 (12) = StaticCall <mem> {AuxCall{runtime.printunlock}} v14
    v16 (12) = SelectN <mem> [0] v15
    v17 (+9) = MakeResult <mem> v16
    v18 (?) = Const64 <int64> [0]
    Ret v17 (9)

    dse [4314 ns]
    memcombine [1300 ns]
    writebarrier [2112 ns]
    lower [16076 ns]
    b2:-
    v1 (?) = InitMem <mem>
    v6 (+7) = InlMark <void> [0] v1
    v7 (12) = MOVQconst <int> [3]
    v8 (+12) = CALLstatic <mem> {AuxCall{runtime.printlock}} v1
    v9 (12) = SelectN <mem> [0] v8
    v11 (12) = CALLstatic <mem> {AuxCall{runtime.printint}} [8] v7 v9
    v12 (12) = SelectN <mem> [0] v11
    v13 (12) = CALLstatic <mem> {AuxCall{runtime.printnl}} v12
    v14 (12) = SelectN <mem> [0] v13
    v15 (12) = CALLstatic <mem> {AuxCall{runtime.printunlock}} v14
    v16 (12) = SelectN <mem> [0] v15
    v17 (+9) = MakeResult <mem> v16
    Ret v17 (9)

    addressing modes [1292 ns]
    late lower [1080 ns]
    lowered deadcode for cse [1528 ns]
    lowered cse [1999 ns]
    elim unread autos [226 ns]
    tighten tuple selectors [1708 ns]
    lowered deadcode [1426 ns]
    checkLower [333 ns]
    late phielim [279 ns]
    late copyelim [765 ns]
    tighten [8510 ns]
    late deadcode [1762 ns]
    critical [467 ns]
    phi tighten [211 ns]
    likelyadjust [992 ns]
    layout [1873 ns]
    schedule [13537 ns]
    b2:-
    v1 (?) = InitMem <mem>
    v7 (12) = MOVQconst <int> [3]
    v6 (+7) = InlMark <void> [0] v1
    v8 (+12) = CALLstatic <mem> {AuxCall{runtime.printlock}} v1
    v9 (12) = SelectN <mem> [0] v8
    v11 (12) = CALLstatic <mem> {AuxCall{runtime.printint}} [8] v7 v9
    v12 (12) = SelectN <mem> [0] v11
    v13 (12) = CALLstatic <mem> {AuxCall{runtime.printnl}} v12
    v14 (12) = SelectN <mem> [0] v13
    v15 (12) = CALLstatic <mem> {AuxCall{runtime.printunlock}} v14
    v16 (12) = SelectN <mem> [0] v15
    v17 (+9) = MakeResult <mem> v16
    Ret v17 (9)

    late nilcheck [1632 ns]
    flagalloc [3800 ns]
    regalloc [32721 ns]
    b2:-
    v1 (?) = InitMem <mem>
    v6 (+7) = InlMark <void> [0] v1
    v8 (+12) = CALLstatic <mem> {AuxCall{runtime.printlock}} v1
    v9 (12) = SelectN <mem> [0] v8
    v18 (12) = MOVQconst <int> [3] : AX
    v11 (12) = CALLstatic <mem> {AuxCall{runtime.printint}} [8] v18 v9 : <>
    v12 (12) = SelectN <mem> [0] v11
    v13 (12) = CALLstatic <mem> {AuxCall{runtime.printnl}} v12
    v14 (12) = SelectN <mem> [0] v13
    v15 (12) = CALLstatic <mem> {AuxCall{runtime.printunlock}} v14
    v16 (12) = SelectN <mem> [0] v15
    v17 (+9) = MakeResult <mem> v16
    Ret v17 (9)

    loop rotate [313 ns]
    stackframe [3722 ns]
    trim [403 ns]
    b2:-
    v1 (?) = InitMem <mem>
    v6 (+7) = InlMark <void> [0] v1
    v8 (+12) = CALLstatic <mem> {AuxCall{runtime.printlock}} v1
    v9 (12) = SelectN <mem> [0] v8
    v18 (12) = MOVQconst <int> [3] : AX
    v11 (12) = CALLstatic <mem> {AuxCall{runtime.printint}} [8] v18 v9 : <>
    v12 (12) = SelectN <mem> [0] v11
    v13 (12) = CALLstatic <mem> {AuxCall{runtime.printnl}} v12
    v14 (12) = SelectN <mem> [0] v13
    v15 (12) = CALLstatic <mem> {AuxCall{runtime.printunlock}} v14
    v16 (12) = SelectN <mem> [0] v15
    v17 (+9) = MakeResult <mem> v16
    Ret v17 (9)

    genssa
    # /app/public/buildbox/08382b4d-0015-11ef-9e73-0242ac16000d/main.go
        00000 (3) TEXT main.main(SB), ABIInternal
        00001 (3) FUNCDATA $0, gclocals·g2BeySu+wFnoycgXfElmcg==(SB)
        00002 (3) FUNCDATA $1, gclocals·g2BeySu+wFnoycgXfElmcg==(SB)
    v6  00003 (7) XCHGL AX, AX
    v8  00004 (+12) PCDATA $1, $0
    v8  00005 (+12) CALL runtime.printlock(SB)
    v18 00006 (12) MOVL $3, AX
    v11 00007 (12) CALL runtime.printint(SB)
    v13 00008 (12) CALL runtime.printnl(SB)
    v15 00009 (12) CALL runtime.printunlock(SB)
    b2  00010 (9) RET
        00011 (?) END
    ```
</details>

### SSA для add()
<details>
    ```console
    add help darkmode 
    sources
    AST
    buildssa-enter
    buildssa-body
    . BLOCK-init
    . . AS tc(1) # main.go:12:9
    . . . NAME-main..autotmp_2 esc(N) Class:PAUTO Offset:0 AutoTemp OnStack Used int tc(1) # main.go:12:9
    . . . ADD int tc(1) # main.go:12:12
    . . . . NAME-main.a esc(no) Class:PPARAM Offset:0 OnStack Used int tc(1) # main.go:11:10
    . . . . NAME-main.b esc(no) Class:PPARAM Offset:0 OnStack Used int tc(1) # main.go:11:13
    . BLOCK tc(1) # main.go:12:9
    . BLOCK-List
    . . CALLFUNC Walked tc(1) # main.go:12:9
    . . . NAME-runtime.printlock Class:PFUNC Offset:0 Used FUNC-func() tc(1)
    . . CALLFUNC Walked tc(1) # main.go:12:9
    . . . NAME-runtime.printint Class:PFUNC Offset:0 Used FUNC-func(int64) tc(1)
    . . CALLFUNC-Args
    . . . CONVNOP int64 tc(1) # main.go:12:9
    . . . . NAME-main..autotmp_2 esc(N) Class:PAUTO Offset:0 AutoTemp OnStack Used int tc(1) # main.go:12:9
    . . CALLFUNC Walked tc(1) # main.go:12:9
    . . . NAME-runtime.printnl Class:PFUNC Offset:0 Used FUNC-func() tc(1)
    . . CALLFUNC Walked tc(1) # main.go:12:9
    . . . NAME-runtime.printunlock Class:PFUNC Offset:0 Used FUNC-func() tc(1)
    buildssa-exit
    
    before insert phis
    b1:-
    v1 (?) = InitMem <mem>
    v2 (?) = SP <uintptr>
    v3 (?) = SB <uintptr>
    v4 (?) = LocalAddr <*int> {a} v2 v1
    v5 (?) = LocalAddr <*int> {b} v2 v1
    v6 (11) = Arg <int> {a} (a[int])
    v7 (11) = Arg <int> {b} (b[int])
    v8 (12) = Add64 <int> v6 v7
    v9 (12) = StaticLECall <mem> {AuxCall{runtime.printlock}} v1
    v10 (12) = SelectN <mem> [0] v9
    v11 (12) = Copy <int64> v8
    v12 (12) = StaticLECall <mem> {AuxCall{runtime.printint}} [8] v11 v10
    v13 (12) = SelectN <mem> [0] v12
    v14 (12) = StaticLECall <mem> {AuxCall{runtime.printnl}} v13
    v15 (12) = SelectN <mem> [0] v14
    v16 (12) = StaticLECall <mem> {AuxCall{runtime.printunlock}} v15
    v17 (12) = SelectN <mem> [0] v16
    v18 (13) = MakeResult <mem> v17
    Ret v18 (13)
    name a[int]: v6
    name b[int]: v7

    start
    number lines [4811 ns]
    b1:-
    v1 (?) = InitMem <mem>
    v2 (?) = SP <uintptr>
    v3 (?) = SB <uintptr>
    v4 (?) = LocalAddr <*int> {a} v2 v1
    v5 (?) = LocalAddr <*int> {b} v2 v1
    v6 (11) = Arg <int> {a} (a[int])
    v7 (11) = Arg <int> {b} (b[int])
    v8 (+12) = Add64 <int> v6 v7
    v9 (12) = StaticLECall <mem> {AuxCall{runtime.printlock}} v1
    v10 (12) = SelectN <mem> [0] v9
    v11 (12) = Copy <int64> v8
    v12 (12) = StaticLECall <mem> {AuxCall{runtime.printint}} [8] v11 v10
    v13 (12) = SelectN <mem> [0] v12
    v14 (12) = StaticLECall <mem> {AuxCall{runtime.printnl}} v13
    v15 (12) = SelectN <mem> [0] v14
    v16 (12) = StaticLECall <mem> {AuxCall{runtime.printunlock}} v15
    v17 (12) = SelectN <mem> [0] v16
    v18 (+13) = MakeResult <mem> v17
    Ret v18 (13)
    name a[int]: v6
    name b[int]: v7

    early phielim [669 ns]
    b1:-
    v1 (?) = InitMem <mem>
    v2 (?) = SP <uintptr>
    v3 (?) = SB <uintptr>
    v4 (?) = LocalAddr <*int> {a} v2 v1
    v5 (?) = LocalAddr <*int> {b} v2 v1
    v6 (11) = Arg <int> {a} (a[int])
    v7 (11) = Arg <int> {b} (b[int])
    v8 (+12) = Add64 <int> v6 v7
    v9 (12) = StaticLECall <mem> {AuxCall{runtime.printlock}} v1
    v10 (12) = SelectN <mem> [0] v9
    v11 (12) = Copy <int64> v8
    v12 (12) = StaticLECall <mem> {AuxCall{runtime.printint}} [8] v8 v10
    v13 (12) = SelectN <mem> [0] v12
    v14 (12) = StaticLECall <mem> {AuxCall{runtime.printnl}} v13
    v15 (12) = SelectN <mem> [0] v14
    v16 (12) = StaticLECall <mem> {AuxCall{runtime.printunlock}} v15
    v17 (12) = SelectN <mem> [0] v16
    v18 (+13) = MakeResult <mem> v17
    Ret v18 (13)
    name a[int]: v6
    name b[int]: v7

    early copyelim [762 ns]
    early deadcode [5159 ns]
    b1:-
    v1 (?) = InitMem <mem>
    v6 (11) = Arg <int> {a} (a[int])
    v7 (11) = Arg <int> {b} (b[int])
    v8 (+12) = Add64 <int> v6 v7
    v9 (12) = StaticLECall <mem> {AuxCall{runtime.printlock}} v1
    v10 (12) = SelectN <mem> [0] v9
    v12 (12) = StaticLECall <mem> {AuxCall{runtime.printint}} [8] v8 v10
    v13 (12) = SelectN <mem> [0] v12
    v14 (12) = StaticLECall <mem> {AuxCall{runtime.printnl}} v13
    v15 (12) = SelectN <mem> [0] v14
    v16 (12) = StaticLECall <mem> {AuxCall{runtime.printunlock}} v15
    v17 (12) = SelectN <mem> [0] v16
    v18 (+13) = MakeResult <mem> v17
    Ret v18 (13)
    name a[int]: v6
    name b[int]: v7

    short circuit [1164 ns]
    decompose user [803 ns]
    pre-opt deadcode [3097 ns]
    opt [10525 ns]
    zero arg cse [2726 ns]
    opt deadcode [2593 ns]
    generic cse [20081 ns]
    phiopt [625 ns]
    gcse deadcode [2698 ns]
    nilcheckelim [16475 ns]
    prove [9708 ns]
    b1:-
    v1 (?) = InitMem <mem>
    v6 (11) = Arg <int> {a} (a[int])
    v7 (11) = Arg <int> {b} (b[int])
    v8 (+12) = Add64 <int> v6 v7
    v9 (12) = StaticLECall <mem> {AuxCall{runtime.printlock}} v1
    v10 (12) = SelectN <mem> [0] v9
    v12 (12) = StaticLECall <mem> {AuxCall{runtime.printint}} [8] v8 v10
    v13 (12) = SelectN <mem> [0] v12
    v14 (12) = StaticLECall <mem> {AuxCall{runtime.printnl}} v13
    v15 (12) = SelectN <mem> [0] v14
    v16 (12) = StaticLECall <mem> {AuxCall{runtime.printunlock}} v15
    v17 (12) = SelectN <mem> [0] v16
    v18 (+13) = MakeResult <mem> v17
    v11 (?) = Const64 <int64> [0]
    Ret v18 (13)
    name a[int]: v6
    name b[int]: v7

    early fuse [929 ns]
    decompose builtin [4806 ns]
    expand calls [30302 ns]
    b1:-
    v1 (?) = InitMem <mem>
    v6 (11) = ArgIntReg <int> {a+0} [0] (a[int])
    v7 (11) = ArgIntReg <int> {b+0} [1] (b[int])
    v8 (+12) = Add64 <int> v6 v7
    v9 (12) = StaticCall <mem> {AuxCall{runtime.printlock}} v1
    v10 (12) = SelectN <mem> [0] v9
    v12 (12) = StaticCall <mem> {AuxCall{runtime.printint}} [8] v8 v10
    v13 (12) = SelectN <mem> [0] v12
    v14 (12) = StaticCall <mem> {AuxCall{runtime.printnl}} v13
    v15 (12) = SelectN <mem> [0] v14
    v16 (12) = StaticCall <mem> {AuxCall{runtime.printunlock}} v15
    v17 (12) = SelectN <mem> [0] v16
    v18 (+13) = MakeResult <mem> v17
    v11 (?) = Const64 <int64> [0]
    v5 (?) = SB <uintptr>
    v4 (?) = SP <uintptr>
    Ret v18 (13)
    name a[int]: v6
    name b[int]: v7

    softfloat [513 ns]
    late opt [7553 ns]
    b1:-
    v1 (?) = InitMem <mem>
    v6 (11) = ArgIntReg <int> {a+0} [0] (a[int])
    v7 (11) = ArgIntReg <int> {b+0} [1] (b[int])
    v8 (+12) = Add64 <int> v6 v7
    v9 (12) = StaticCall <mem> {AuxCall{runtime.printlock}} v1
    v10 (12) = SelectN <mem> [0] v9
    v12 (12) = StaticCall <mem> {AuxCall{runtime.printint}} [8] v8 v10
    v13 (12) = SelectN <mem> [0] v12
    v14 (12) = StaticCall <mem> {AuxCall{runtime.printnl}} v13
    v15 (12) = SelectN <mem> [0] v14
    v16 (12) = StaticCall <mem> {AuxCall{runtime.printunlock}} v15
    v17 (12) = SelectN <mem> [0] v16
    v18 (+13) = MakeResult <mem> v17
    Ret v18 (13)
    name a[int]: v6
    name b[int]: v7

    dead auto elim [1980 ns]
    generic deadcode [3224 ns]
    check bce [229 ns]
    branchelim [1373 ns]
    late fuse [3886 ns]
    b1:-
    v1 (?) = InitMem <mem>
    v6 (11) = ArgIntReg <int> {a+0} [0] (a[int])
    v7 (11) = ArgIntReg <int> {b+0} [1] (b[int])
    v8 (+12) = Add64 <int> v6 v7
    v9 (12) = StaticCall <mem> {AuxCall{runtime.printlock}} v1
    v10 (12) = SelectN <mem> [0] v9
    v12 (12) = StaticCall <mem> {AuxCall{runtime.printint}} [8] v8 v10
    v13 (12) = SelectN <mem> [0] v12
    v14 (12) = StaticCall <mem> {AuxCall{runtime.printnl}} v13
    v15 (12) = SelectN <mem> [0] v14
    v16 (12) = StaticCall <mem> {AuxCall{runtime.printunlock}} v15
    v17 (12) = SelectN <mem> [0] v16
    v18 (+13) = MakeResult <mem> v17
    v4 (?) = Const64 <int64> [0]
    Ret v18 (13)
    name a[int]: v6
    name b[int]: v7

    dse [4151 ns]
    memcombine [1761 ns]
    writebarrier [2856 ns]
    lower [7862 ns]
    b1:-
    v1 (?) = InitMem <mem>
    v6 (11) = ArgIntReg <int> {a+0} [0] (a[int])
    v7 (11) = ArgIntReg <int> {b+0} [1] (b[int])
    v8 (+12) = ADDQ <int> v6 v7
    v9 (12) = CALLstatic <mem> {AuxCall{runtime.printlock}} v1
    v10 (12) = SelectN <mem> [0] v9
    v12 (12) = CALLstatic <mem> {AuxCall{runtime.printint}} [8] v8 v10
    v13 (12) = SelectN <mem> [0] v12
    v14 (12) = CALLstatic <mem> {AuxCall{runtime.printnl}} v13
    v15 (12) = SelectN <mem> [0] v14
    v16 (12) = CALLstatic <mem> {AuxCall{runtime.printunlock}} v15
    v17 (12) = SelectN <mem> [0] v16
    v18 (+13) = MakeResult <mem> v17
    Ret v18 (13)
    name a[int]: v6
    name b[int]: v7

    addressing modes [1627 ns]
    late lower [1813 ns]
    lowered deadcode for cse [3165 ns]
    lowered cse [4210 ns]
    elim unread autos [272 ns]
    tighten tuple selectors [1601 ns]
    lowered deadcode [1960 ns]
    checkLower [422 ns]
    late phielim [522 ns]
    late copyelim [688 ns]
    tighten [9121 ns]
    late deadcode [2783 ns]
    critical [614 ns]
    phi tighten [325 ns]
    likelyadjust [1258 ns]
    layout [2165 ns]
    schedule [21343 ns]
    b1:-
    v6 (11) = ArgIntReg <int> {a+0} [0] (a[int])
    v7 (11) = ArgIntReg <int> {b+0} [1] (b[int])
    v1 (?) = InitMem <mem>
    v9 (12) = CALLstatic <mem> {AuxCall{runtime.printlock}} v1
    v10 (12) = SelectN <mem> [0] v9
    v8 (+12) = ADDQ <int> v6 v7
    v12 (12) = CALLstatic <mem> {AuxCall{runtime.printint}} [8] v8 v10
    v13 (12) = SelectN <mem> [0] v12
    v14 (12) = CALLstatic <mem> {AuxCall{runtime.printnl}} v13
    v15 (12) = SelectN <mem> [0] v14
    v16 (12) = CALLstatic <mem> {AuxCall{runtime.printunlock}} v15
    v17 (12) = SelectN <mem> [0] v16
    v18 (+13) = MakeResult <mem> v17
    Ret v18 (13)
    name a[int]: v6
    name b[int]: v7

    late nilcheck [1844 ns]
    flagalloc [3524 ns]
    regalloc [44412 ns]
    b1:-
    v6 (11) = ArgIntReg <int> {a+0} [0] : AX (a[int])
    v7 (11) = ArgIntReg <int> {b+0} [1] : BX (b[int])
    v11 (11) = StoreReg <int> v6 : a[int]
    v4 (11) = StoreReg <int> v7 : b[int]
    v1 (?) = InitMem <mem>
    v9 (12) = CALLstatic <mem> {AuxCall{runtime.printlock}} v1
    v10 (12) = SelectN <mem> [0] v9
    v5 (12) = LoadReg <int> v4 : AX
    v3 (12) = LoadReg <int> v11 : CX
    v8 (+12) = ADDQ <int> v3 v5 : AX
    v12 (12) = CALLstatic <mem> {AuxCall{runtime.printint}} [8] v8 v10 : <>
    v13 (12) = SelectN <mem> [0] v12
    v14 (12) = CALLstatic <mem> {AuxCall{runtime.printnl}} v13
    v15 (12) = SelectN <mem> [0] v14
    v16 (12) = CALLstatic <mem> {AuxCall{runtime.printunlock}} v15
    v17 (12) = SelectN <mem> [0] v16
    v18 (+13) = MakeResult <mem> v17
    Ret v18 (13)
    name a[int]: v6
    name b[int]: v7

    loop rotate [348 ns]
    stackframe [3835 ns]
    trim [396 ns]
    b1:-
    v6 (11) = ArgIntReg <int> {a+0} [0] : AX (a[int])
    v7 (11) = ArgIntReg <int> {b+0} [1] : BX (b[int])
    v11 (11) = StoreReg <int> v6 : a[int]
    v4 (11) = StoreReg <int> v7 : b[int]
    v1 (?) = InitMem <mem>
    v9 (12) = CALLstatic <mem> {AuxCall{runtime.printlock}} v1
    v10 (12) = SelectN <mem> [0] v9
    v5 (12) = LoadReg <int> v4 : AX
    v3 (12) = LoadReg <int> v11 : CX
    v8 (+12) = ADDQ <int> v3 v5 : AX
    v12 (12) = CALLstatic <mem> {AuxCall{runtime.printint}} [8] v8 v10 : <>
    v13 (12) = SelectN <mem> [0] v12
    v14 (12) = CALLstatic <mem> {AuxCall{runtime.printnl}} v13
    v15 (12) = SelectN <mem> [0] v14
    v16 (12) = CALLstatic <mem> {AuxCall{runtime.printunlock}} v15
    v17 (12) = SelectN <mem> [0] v16
    v18 (+13) = MakeResult <mem> v17
    Ret v18 (13)
    name a[int]: v6
    name b[int]: v7

    genssa
    # /app/public/buildbox/abbe7c00-0017-11ef-9e73-0242ac16000d/main.go
        00000 (11) TEXT main.add(SB), ABIInternal
        00001 (11) FUNCDATA $0, gclocals·g2BeySu+wFnoycgXfElmcg==(SB)
        00002 (11) FUNCDATA $1, gclocals·g2BeySu+wFnoycgXfElmcg==(SB)
        00003 (11) FUNCDATA $5, main.add.arginfo1(SB)
        00004 (11) FUNCDATA $6, main.add.argliveinfo(SB)
    b1  00005 (11) PCDATA $3, $1
    v11 00006 (11) MOVQ AX, main.a(SP)
    v4  00007 (11) MOVQ BX, main.b+8(SP)
    v4  00008 (11) PCDATA $3, $-1
    v9  00009 (+12) PCDATA $1, $0
    v9  00010 (+12) CALL runtime.printlock(SB)
    v5  00011 (12) MOVQ main.b+8(SP), AX
    v3  00012 (12) MOVQ main.a(SP), CX
    v8  00013 (12) ADDQ CX, AX
    v12 00014 (12) CALL runtime.printint(SB)
    v14 00015 (12) CALL runtime.printnl(SB)
    v16 00016 (12) CALL runtime.printunlock(SB)
    b1  00017 (13) RET
        00018 (?) END
    ```
</details>